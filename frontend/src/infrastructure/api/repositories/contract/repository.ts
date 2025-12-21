import api, { Api } from "../../api";
import Contract from "@model/contract/entity";
import {Format, ImportResult, ListResult} from "../../dto";
import Person from "@model/person/model";

export enum ContractRole {
    SIGNATORY = "signatory",
    COUNTERPARTY = "counterparty",
    BENEFICIARY = "beneficiary",
    WITNESS = "witness",
}

export interface CreateContractDTO {
    code: string
    title: string

    description?: string
    startDate?: string
    endDate?: string
}

export interface AddPersonDTO {
    contractId: string,
    personId: string,
    role: ContractRole
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

    async getPersonsByContractId(id: string): Promise<Person[]> {
        return this.api.get(`/contract/${id}/persons`)
    }

    async addPerson(dto: AddPersonDTO): Promise<any> {
        return this.api.post("/contract/person", dto)
    }

    async removePerson(contractId: string, personId: string): Promise<boolean> {
        try {
            await this.api.delete(`/contract/${contractId}/person/${personId}`)
            return true
        } catch {
            return false
        }
    }

    async getContractsByPersonId(id: string): Promise<Contract[]> {
        return this.api.get(`/contract/person/${id}`)
    }

    async export(format: Format): Promise<void> {
        const blob = await api.get<Blob>(
            `/contract/export?format=${format}`,
            { responseType: "blob" }
        );

        const url = URL.createObjectURL(blob);
        const a = document.createElement("a");

        a.href = url;
        a.download = `export.${format}`;
        a.click();

        URL.revokeObjectURL(url);
    }

    async import(file: File): Promise<ImportResult> {
        const formData = new FormData();
        formData.append("file", file);

        return api.post(
            "/contract/import",
            formData,
            { isFormData: true }
        );
    }
}

const contractRepository = new ContractRepository()

export default contractRepository;