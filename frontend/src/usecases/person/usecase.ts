import personRepository, {CreatePersonDTO, CreatePersonResponse, PersonRepository} from "@repository/person/repository";
import {ListResult} from "../../infrastructure/api/dto";
import Person from "@model/person/model";

export class PersonUsecase {
    private personRepository: PersonRepository = personRepository;

    async get(id: string): Promise<Person> {
        return this.personRepository.get(id);
    }

    async create(dto: CreatePersonDTO): Promise<boolean> {
        return this.personRepository.create(dto);
    }

    async update(dto: any, id: string): Promise<boolean> {
        return this.personRepository.update(dto, id);
    }

    async list(filters: string): Promise<ListResult<Person>> {
        return this.personRepository.list(filters);
    }
}

const personUsecase = new PersonUsecase();

export default personUsecase;