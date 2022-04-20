export type Route = {
  id: number;
  selected: boolean;
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  description: string;
  url_param: UrlParam | null;
  query_params: QueryParam[];
};

export type ParamType = "string" | "integer" | "boolean";
export type UrlParam = {
  name: string;
  type: ParamType;
};

export type QueryParam = {
  name: string;
  type: ParamType;
};
