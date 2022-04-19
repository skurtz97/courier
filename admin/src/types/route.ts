export type Route = {
  id: number;
  selected: boolean;
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  description: string;
  url_param: UrlParam | null;
  query_params: QueryParam[] | null;
};

export type UrlParam = {
  name: string;
  type: "string" | "integer" | "boolean";
};

export type QueryParam = {
  name: string;
  type: "string" | "integer" | "boolean";
};
