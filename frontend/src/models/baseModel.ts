interface BaseModel {
    id: string;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string | null;
}

export default BaseModel;