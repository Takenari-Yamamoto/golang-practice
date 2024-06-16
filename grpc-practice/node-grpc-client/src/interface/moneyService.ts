import { PromiseClient, createPromiseClient } from "@connectrpc/connect";
import { MoneyService } from "../../generated/proto/money_connect";
import { vi } from "vitest";

export type IMoneyServiceClient = PromiseClient<typeof MoneyService>;
