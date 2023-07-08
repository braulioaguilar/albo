import { Controller, Get, HttpStatus, Param, Res } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/api/characters/:character')
  async getCharacters(@Res() response, @Param('character') character: string) {
    try {
      const characters = await this.appService.findOne(character);
      return response.status(HttpStatus.OK).json({
        last_sync: characters.createdAt,
        characters: characters.characters
      });
    } catch (error) {
      return response.status(error.status).json(error.response);
    }
  }
}
