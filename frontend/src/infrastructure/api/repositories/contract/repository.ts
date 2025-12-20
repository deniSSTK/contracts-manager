import api, { Api } from "../../api";
import Contract from "@model/contract/entity";
import {ListResult} from "../../dto";

export interface CreateContractDTO {
    code: string
    title: string

    description?: string
    startDate?: string
    endDate?: string
}

export class ContractRepository {
    private readonly api: Api = api;

    async get(id: string): Promise<Contract> {
        return this.api.get(`/contract/${id}`)
    }

    async create(dto: CreateContractDTO): Promise<boolean> {
        try {
            await this.api.post("/contract/", dto);
            return true;
        } catch {
            return false;
        }
    }

    async update(dto: any, id: string): Promise<boolean> {
        try {
            await this.api.put(`/contract/${id}`, dto);
            return true
        } catch {
            return false
        }
    }

    async list(filters: string): Promise<ListResult<Contract>> {
        return this.api.get("/contract/?" + filters)
    }
}

const contractRepository = new ContractRepository()

export default contractRepository;