import { MockClient } from "../utils/mock";

export type UserServiceClient = {
  getUserById: (request: { id: string }) => Promise<{
    id: string;
    name: string;
  }>;
};

// Mock の UserServiceClient を作成する
export class MockUserServiceClient extends MockClient<UserServiceClient> {
  constructor() {
    super();
  }
}
