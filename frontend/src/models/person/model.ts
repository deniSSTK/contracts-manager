import BaseModel from "@model/baseModel";

export type PersonType = 'individual' | 'entity';

interface Person extends BaseModel {
    type: PersonType;
    name: string;
    code: string;
    email?: string | null;
    phone?: string | null;
}

export default Person;