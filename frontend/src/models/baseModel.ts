interface BaseModel {
    id: string;
    cratedAt: string;
    updatedAt: string;
    deletedAt?: string | null;
}

export default BaseModel;