import personRepository, {CreatePersonDTO, CreatePersonResponse, PersonRepository} from "@repository/person/repository";
import {ListResult} from "../../infrastructure/api/dto";
import Person from "@model/person/model";

export class PersonUsecase {
    private personRepository: PersonRepository = personRepository;

    async list(filters: string): Promise<ListResult<Person>> {
        return this.personRepository.list(filters);
    }

    async create(dto: CreatePersonDTO): Promise<CreatePersonResponse> {
        return this.personRepository.create(dto);
    }
}

const personUsecase = new PersonUsecase();

export default personUsecase;