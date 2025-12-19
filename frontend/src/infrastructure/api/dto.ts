import BaseModel from "@model/baseModel";

export interface ListResult<T extends BaseModel> {
    data: T[];
    page: number;
    limit: number;
    total: number;
}