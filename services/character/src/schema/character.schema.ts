import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose";
import { CharacterDocument } from "../interface/character.interface";

@Schema()
export class Character {
    @Prop()
    createdAt: string;

    @Prop()
    name: string;

    @Prop()
    characters: Array<CharacterDocument>;
}

export const CharacterSchema = SchemaFactory.createForClass(Character);