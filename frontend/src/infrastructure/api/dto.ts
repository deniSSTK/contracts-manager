import BaseModel from "@model/baseModel";

export type Format = 'csv' | 'json'

export interface ListResult<T extends BaseModel> {
    data: T[];
    page: number;
    limit: number;
    total: number;
}

export interface ImportResult {
    imported: number;
    errors: string[];
}