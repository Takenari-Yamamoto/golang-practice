import { HelloServiceClient } from "../interface/helloServiceClient";
import { UserServiceClient } from "../interface/userService";

export class UseCase {
  private helloServiceClient: HelloServiceClient;
  private userServiceClient: UserServiceClient;

  constructor({
    helloService,
    userService,
  }: {
    helloService: HelloServiceClient;
    userService: UserServiceClient;
  }) {
    this.helloServiceClient = helloService;
    this.userServiceClient = userService;
  }

  async sayHello(id: string): Promise<string> {
    const user = await this.userServiceClient.getUserById({ id });
    const response = await this.helloServiceClient.sayHello({ id: user.id });
    return response.message;
  }
}
