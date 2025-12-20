import contractRepository, {ContractRepository, CreateContractDTO} from "@repository/contract/repository";
import Contract from "@model/contract/entity";
import {ListResult} from "../../infrastructure/api/dto";

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
}

const contractUsecase = new ContractUsecase();

export default contractUsecase;