import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { CharacterService } from './character/character.service';
import { CharacterController } from './character/character.controller';

@Module({
  imports: [MongooseModule.forRoot('mongodb://localhost:27017/albo')],
  controllers: [AppController, CharacterController],
  providers: [AppService, CharacterService],
})
export class AppModule {}
