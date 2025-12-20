import BaseModel from "@model/baseModel";

interface Contract extends BaseModel {
    code: string
    title: string

    description?: string
    startDate?: string
    endDate?: string
}

export default Contract;