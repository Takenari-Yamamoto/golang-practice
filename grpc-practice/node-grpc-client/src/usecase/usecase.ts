import { BookServiceClient } from "../interface/bookService";
import { UserServiceClient } from "../interface/userService";

export class UseCase {
  private userServiceClient: UserServiceClient;
  private bookServiceClient: BookServiceClient;

  constructor({
    userService,
    bookService,
  }: {
    userService: UserServiceClient;
    bookService: BookServiceClient;
  }) {
    this.userServiceClient = userService;
    this.bookServiceClient = bookService;
  }

  async getUserById(id: string) {
    const user = await this.userServiceClient.getUserById({ id });
    const book = await this.bookServiceClient.listBooksByUserId({
      userId: user.id,
    });

    return {
      ...user,
      book,
    };
  }
}
