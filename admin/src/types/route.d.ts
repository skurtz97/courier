export type RequestMethod = "GET" | "POST" | "PUT" | "DELETE";

export type Route = {
	method: RequestMethod;
	path: string;
	description: string;
	url_param: UrlParam;
	query_params: QueryParam[] | null;
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
