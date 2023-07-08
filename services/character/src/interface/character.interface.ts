import { Document } from "mongoose"

export interface CharacterDocument extends Document {
  readonly character: string

  readonly comics: Array<string>
}
