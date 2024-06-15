import { MockClient } from "../utils/mock";

export type HelloServiceClient = {
  readonly sayHello: (request: { id: string }) => Promise<{
    message: string;
  }>;
};

export class MockHelloServiceClient extends MockClient<HelloServiceClient> {
  constructor() {
    super();
  }
}
