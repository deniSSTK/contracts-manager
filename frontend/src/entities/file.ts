import {Format, ImportResult} from "../infrastructure/api/dto";
import personUsecase from "@usecase/person/usecase";
import contractUsecase from "@usecase/contract/usecase";

interface FileUsecase {
    export(format: Format): Promise<void>
    import(file: File): Promise<ImportResult>
}

const fileEntities: Record<string, FileUsecase> = {
    person: personUsecase,
    contract: contractUsecase
}

export default fileEntities;