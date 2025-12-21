import personRepository, {CreatePersonDTO, CreatePersonResponse, PersonRepository} from "@repository/person/repository";
import {Format, ImportResult, ListResult} from "../../infrastructure/api/dto";
import Person from "@model/person/model";
import Contract from "@model/contract/entity";

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

    async export(format: Format): Promise<void> {
        return this.personRepository.export(format)
    }

    async import(file: File): Promise<ImportResult> {
        return this.personRepository.import(file)
    }
}

const personUsecase = new PersonUsecase();

export default personUsecase;