import personUsecase from "@usecase/person/usecase"
import contractUsecase from "@usecase/contract/usecase"

import { ListResult } from "../infrastructure/api/dto"
import BaseModel from "@model/baseModel"
import Person from "@model/person/model"
import router from "@app/router/routes";
import {RouteName} from "@app/router/types";
import authUsecase from "@usecase/auth/usecase";
import Contract from "@model/contract/entity";
import {User} from "@model/user/model";

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
    callback: (row: T) => void
}

interface EntityColumn<T = any> {
    key: string
    label: string

    optional?: boolean;

    actions?: EntityAction<T>[]
}

function infoAction<T extends BaseModel>(entity: string): EntityAction<T> {
    return {
        label: "Info",
        callback: async (row) =>
            await router.push({
                name: RouteName.ADMIN_PANEL_INFO,
                params: {
                    entity,
                    entityId: row.id,
                },
            }),
    }
}

const entityRegistry: Record<string, EntityConfig> = {
    person: {
        usecase: personUsecase,
        columns: [
            { key: "code", label: "Code" },
            { key: "name", label: "Name" },
            { key: "type", label: "Type" },
            { key: "id", label: "Id" },
            { key: "email", label: "Email", optional: true },
            { key: "phone", label: "Phone", optional: true },
            {
                key: "actions",
                label: "actions",
                actions: [infoAction<Person>("person")],
            },
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
            { key: "code", label: "Code" },
            { key: "title", label: "Title" },
            { key: "description", label: "Description", optional: true },
            { key: "startDate", label: "Start Date", optional: true },
            { key: "endDate", label: "End Date", optional: true },
            {
                key: "actions",
                label: "actions",
                actions: [infoAction<Contract>("contract")],
            },
        ],
        filters: [
            { key: "code", placeholder: "Code" },
            { key: "title", placeholder: "Title" },
            { key: "description", placeholder: "Description" },
        ],
    },

    user: {
        usecase: authUsecase,
        columns: [
            { key: "username", label: "Username" },
            { key: "email", label: "Email" },
            { key: "type", label: "Type" },
            { key: "personId", label: "PersonID", optional: true },
            {
                key: "actions",
                label: "actions",
                actions: [infoAction<User>("user")],
            },
        ],
        filters: [
            { key: "username", placeholder: "Username" },
            { key: "email", placeholder: "Email" },
            { key: "type", placeholder: "Type" },
        ]
    }
}
export default entityRegistry;