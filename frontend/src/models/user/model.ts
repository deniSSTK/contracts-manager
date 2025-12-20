import IBaseModel from "../baseModel";
import Person from "../person/model";
import {UserType} from "@store/auth/store";

export interface User extends IBaseModel {
    username: string;
    email: string;
    type: UserType;
    personID?: string | null;
    person?: Person | null;
}
