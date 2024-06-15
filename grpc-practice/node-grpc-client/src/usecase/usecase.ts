import { HelloServiceClient } from "../interface/helloServiceClient";

export class UseCase {
  private helloServiceClient: HelloServiceClient;

  constructor({ helloService }: { helloService: HelloServiceClient }) {
    this.helloServiceClient = helloService;
  }

  async sayHello(id: string): Promise<string> {
    const { message } = await this.helloServiceClient.sayHello({ id });
    return message;
  }
}
