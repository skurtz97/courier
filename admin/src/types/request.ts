export type Request = {
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  body: string;
};
