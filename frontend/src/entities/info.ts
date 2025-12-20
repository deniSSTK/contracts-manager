import BaseModel from '@model/baseModel';
import personUsecase from '@usecase/person/usecase';

interface InfoRow {
    key: string,
    label: string,

    values?: string[],
    optional?: boolean,
    canChange?: boolean,
}

interface InfoUsecase<T extends BaseModel = BaseModel> {
    get: (id: string) => Promise<T>
}

interface Info {
    usecase: InfoUsecase,
    rows: InfoRow[],
}

const BaseModelRows: InfoRow[] = [
    {
        key: 'id',
        label: 'Id',
    },
    {
        key: 'createdAt',
        label: 'Created At',
    },
    {
        key: 'updatedAt',
        label: 'Last Update At',
    },
]

const infoEntities: Record<string, Info> = {
    person: {
        usecase: personUsecase,
        rows: [
            ...BaseModelRows,
            {
                key: 'name',
                label: 'Name',
                canChange: true,
            },
            {
                key: 'type',
                label: 'Type',
                values: ['individual', 'entity'],
                canChange: true,
            },
            {
                key: 'code',
                label: 'Code',
                canChange: true,
            },
            {
                key: 'email',
                label: 'Email',
                optional: true,
                canChange: true,
            },
            {
                key: 'phone',
                label: 'Phone',
                optional: true,
                canChange: true,
            },
        ]
    },
}

export default infoEntities;