import api, { Api } from "../../api";
import {ListResult} from "../../dto";
import Person from "@model/person/model";

export interface CreatePersonDTO {
    type: "individual" | "entity" | "";
    name: string;
    code: string;
    email?: string | null;
    phone?: string | null;
}

export interface CreatePersonResponse {
    personId: string
}

export class PersonRepository {
    private readonly api: Api = api;

    async list(filters: string): Promise<ListResult<Person>> {
        return this.api.get("/person/?" + filters)
    }

    async create(dto: CreatePersonDTO): Promise<CreatePersonResponse> {
        return this.api.post("/person/", dto)
    }
}

const personRepository = new PersonRepository();

export default personRepository;