import { ForbiddenException, BadRequestException, Injectable, OnModuleInit, NotFoundException, InternalServerErrorException } from '@nestjs/common';
import { forEach } from 'lodash';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Character } from 'src/schema/character.schema';

// TODO: move to env file
const MARLVEL_API = 'http://gateway.marvel.com/v1/public/';
const TS = '1688497908';
const API_KEY = '9b2f073ee571a47da65a782a644cfffa';
const HASH = '53c019e897a0467afba1614fdd342df9';

@Injectable()
export class AppService implements OnModuleInit {
  constructor(@InjectModel('Character') private characterModel: Model<Character>) { }

  async onModuleInit() {
    await this.sync()
  }

  async findOne(character: string) {
    const existingStudent = await this.characterModel.findOne({name: character}).exec();
    if (!existingStudent) {
      throw new NotFoundException(`Character #${character} not found`);
    }
    return existingStudent;
  }

  async create(name: string, characters: Array<Character>) {
    try {
      const createdUrl = new this.characterModel({
        createdAt: new Date(),
        name: name,
        characters: characters,
      });
      await createdUrl.save();      
    } catch (error) {
      throw new InternalServerErrorException(error);
    }
  }

  async sync() {
    console.log("Sync init...")
    try {
      const characterList = [
        {
          nickname: 'ironman',
          name: 'Iron Man'
        }, {
          nickname: 'capamerica',
          name: 'Captain America'
        }
      ]

      let comics = [];
      for (const character of characterList) {
        const characterRequest = await fetch(`${MARLVEL_API}/characters?apikey=${API_KEY}&ts=${TS}&hash=${HASH}&name=${character.name}`);
        const characterResponse = await characterRequest.json();
        const characterRes = characterResponse.data.results;

        if (characterRes.length == 0) {
          throw new BadRequestException('character not found')
        }

        const characterItem = characterRes[0];
        const comicsRequest = await fetch(`${MARLVEL_API}/characters/${characterItem.id}/comics?apikey=${API_KEY}&ts=${TS}&hash=${HASH}`);
        console.log(`Getting comics for ${character.name}...`)
        const comicsResponse = await comicsRequest.json();
        const results = comicsResponse.data.results;

        forEach(results, (comic) => {
          const { title, characters } = comic;
          forEach(characters.items, (item) => {
            comics.push({
              comic: title,
              character: item.name
            })
          })
        })

        // Save each item
        const findCharacter = (characters = [], characterName) => {
          return characters.findIndex(c => c.character === characterName)
        }
  
        const newComics = comics.reduce((characters, comic) => {
          const index = findCharacter(characters, comic.character)
          if (index === -1) {
            return [...characters, { character: comic.character, comics: [comic.comic] }]
          }
  
          characters[index].comics.push(comic.comic)
          return characters
        }, []);
  
        this.create(character.nickname, newComics);
      }
      console.log("Sync done!");
    } catch (error) {
      throw new ForbiddenException(error);
    }
  }
}
