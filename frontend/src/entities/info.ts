import personUsecase from '@usecase/person/usecase'
import { CreatePersonDTO } from '@repository/person/repository'
import authUsecase from "@usecase/auth/usecase";
import contractUsecase from "@usecase/contract/usecase";
import {CreateContractDTO} from "@repository/contract/repository";

interface InfoRow {
    key: string
    label: string

    type?: 'time',

    values?: string[]

    min?: number;
    max?: number;

    optional?: boolean
    canChange?: boolean
}

interface InfoBase {
    usecase: {
        get: (id: string) => Promise<any>
        update: (dto: any, id: string) => Promise<boolean>
    }
    rows: InfoRow[]
}

interface InfoWithCreate<CreateDto> extends InfoBase {
    canCreate: true
    usecase: InfoBase['usecase'] & {
        create: (dto: CreateDto) => Promise<boolean>
    }
    createDto: (form: Record<string, any>) => CreateDto
}

interface InfoWithoutCreate extends InfoBase {
    canCreate: false
}

export type Info<CreateDto = any> =
    | InfoWithCreate<CreateDto>
    | InfoWithoutCreate

const BaseModelRows: InfoRow[] = [
    { key: 'id', label: 'Id' },
    { key: 'createdAt', label: 'Created At', type: 'time' },
    { key: 'updatedAt', label: 'Last Update At', type: 'time' },
]

const infoEntities: Record<string, Info> = {
    person: {
        canCreate: true,
        usecase: personUsecase,
        rows: [
            ...BaseModelRows,
            {key: 'name', label: 'Name', canChange: true, min: 2, max: 255},
            {key: 'type', label: 'Type', values: ['individual', 'entity'], canChange: true},
            {key: 'code', label: 'Code', canChange: true, min: 3, max: 50},
            {key: 'email', label: 'Email', optional: true, canChange: true},
            {key: 'phone', label: 'Phone', optional: true, canChange: true},
        ],
        createDto: (form): CreatePersonDTO => ({
            name: form.name,
            type: form.type,
            code: form.code,
            email: form.email,
            phone: form.phone,
        }),
    },
    user: {
        canCreate: false,
        usecase: authUsecase,
        rows: [
            ...BaseModelRows,
            {key: 'username', label: 'Username', canChange: true, min: 5, max: 50},
            {key: 'email', label: 'Email', canChange: true, max: 100},
            {key: 'type', label: 'Type', values: ["admin", "regular"], canChange: true},
            {key: 'personId', label: 'PersonID', optional: true, canChange: true},
        ]
    },
    contract: {
        canCreate: true,
        usecase: contractUsecase,
        rows: [
            ...BaseModelRows,
            {key: 'code', label: 'Code', canChange: true, min: 1, max: 100},
            {key: 'title', label: 'Title', canChange: true, min: 1, max: 255},
            {key: 'description', label: 'Description', canChange: true, optional: true},
            {key: 'startDate', label: 'Start Date', canChange: true, optional: true, type: 'time'},
            {key: 'endDate', label: 'End Date', canChange: true, optional: true, type: 'time'},
        ],
        createDto: (form): CreateContractDTO => ({
            code: form.code,
            title: form.title,
            description: form.description,
            startDate: form.startDate,
            endDate: form.endDate,
        }),
    },
}

export default infoEntities
