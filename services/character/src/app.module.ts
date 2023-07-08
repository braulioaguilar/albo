import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { CharacterSchema } from './schema/character.schema';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://mongo_albo:27017/albo'),
    MongooseModule.forFeature([{ name: 'Character', schema: CharacterSchema }])
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
