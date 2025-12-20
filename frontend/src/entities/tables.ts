import personUsecase from "@usecase/person/usecase"
import contractUsecase from "@usecase/contract/usecase"
import userUsecase from "@usecase/user/usecase"

import { ListResult } from "../infrastructure/api/dto"
import BaseModel from "@model/baseModel"
import Person from "@model/person/model"
import router, {RouteName} from "@app/router/routes";

export interface EntityFilter {
    key: string
    placeholder: string
}

interface EntityUsecase<T extends BaseModel = BaseModel> {
    list: (filters: string) => Promise<ListResult<T>>
}

interface EntityConfig<T extends BaseModel = BaseModel> {
    usecase: EntityUsecase<T>
    columns: EntityColumn[]
    filters: EntityFilter[]
}

interface EntityAction<T = any> {
    label: string
    callback: (row: T) => void | Promise<void>
}

interface EntityColumn<T = any> {
    key: string
    label: string

    actions?: EntityAction<T>[]
}

export const entityRegistry = {
    person: {
        usecase: personUsecase,
        columns: [
            { key: "code", label: "Code" },
            { key: "name", label: "Name" },
            { key: "type", label: "Type" },
            { key: "id", label: "Id" },
            {
                key: "contracts",
                label: "Check contracts",
                actions: [
                    {
                        label: "Info",
                        callback: async (row: Person) => {
                            await router.push({
                                name: RouteName.ADMIN_PANEL_INFO,
                                params: {
                                    entity: "person",
                                    entityId: row.id,
                                },
                            })
                        },
                    }
                ]
            }
        ],
        filters: [
            { key: "code", placeholder: "Code" },
            { key: "name", placeholder: "Name" },
            { key: "type", placeholder: "Type" },
        ]
    },

    contract: {
        usecase: contractUsecase,
        columns: [
            { key: "title", label: "Title" },
            { key: "status", label: "Status" },
            { key: "number", label: "Number" },
        ],
        filters: [
            { key: "title", placeholder: "Title" },
            { key: "status", placeholder: "Status" },
        ]
    },

    user: {
        usecase: userUsecase,
        columns: [
            { key: "email", label: "Email" },
            { key: "role", label: "Role" },
        ],
        filters: [
            { key: "email", placeholder: "Email" },
        ]
    }
} satisfies Record<string, EntityConfig>