export type Route = {
  id: number;
  selected: boolean;
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  description: string;
};
