import api, { Api } from "../../api";
import {ListResult} from "../../dto";
import Person from "@model/person/model";

export interface CreatePersonDTO {
    type: "individual" | "entity";
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

    async get(id: string): Promise<Person> {
        return this.api.get(`/person/${id}`)
    }

    async create(dto: CreatePersonDTO): Promise<boolean> {
        try {
            await this.api.post("/person/", dto);
            return true;
        } catch {
            return false;
        }
    }

    async update(dto: any, id: string): Promise<boolean> {
        try {
            await this.api.put(`/person/${id}`, dto);
            return true
        } catch {
            return false
        }
    }

    async list(filters: string): Promise<ListResult<Person>> {
        return this.api.get("/person/?" + filters)
    }
}

const personRepository = new PersonRepository();

export default personRepository;