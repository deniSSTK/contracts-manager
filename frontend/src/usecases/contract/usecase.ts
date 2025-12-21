import contractRepository, {AddPersonDTO, ContractRepository, CreateContractDTO} from "@repository/contract/repository";
import Contract from "@model/contract/entity";
import {Format, ImportResult, ListResult} from "../../infrastructure/api/dto";
import Person from "@model/person/model";

class ContractUsecase {
    private contractRepository: ContractRepository = contractRepository;

    async get(id: string): Promise<Contract> {
        return this.contractRepository.get(id);
    }

    async create(dto: CreateContractDTO): Promise<boolean> {
        return this.contractRepository.create(dto);
    }

    async update(dto: any, id: string): Promise<boolean> {
        return this.contractRepository.update(dto, id);
    }

    async list(filters: string): Promise<ListResult<Contract>> {
        return this.contractRepository.list(filters);
    }

    async getPersonsByContractId(id: string): Promise<Person[]> {
        return this.contractRepository.getPersonsByContractId(id)
    }

    async addPerson(dto: AddPersonDTO): Promise<any> {
        return this.contractRepository.addPerson(dto)
    }

    async removePerson(contractId: string, personId: string): Promise<boolean> {
        return this.contractRepository.removePerson(contractId, personId)
    }

    async getContractsByPersonId(id: string): Promise<Contract[]> {
        return this.contractRepository.getContractsByPersonId(id)
    }

    async export(format: Format): Promise<void> {
        return this.contractRepository.export(format)
    }

    async import(file: File): Promise<ImportResult> {
        return this.contractRepository.import(file)
    }
}

const contractUsecase = new ContractUsecase();

export default contractUsecase;