/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import { useInfiniteQuery, useQuery } from "@tanstack/react-query";
import type {
  DataTag,
  DefinedInitialDataOptions,
  DefinedUseInfiniteQueryResult,
  DefinedUseQueryResult,
  InfiniteData,
  QueryFunction,
  QueryKey,
  UndefinedInitialDataOptions,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseQueryOptions,
  UseQueryResult,
} from "@tanstack/react-query";

import { customInstance } from "../../mutator/custom-instance";

type SecondParameter<T extends (...args: never) => unknown> = Parameters<T>[1];

/**
 * Render the home page
 * @summary home admin
 */
export type adminHome1Response301 = {
  data: void;
  status: 301;
};

export type adminHome1ResponseComposite = adminHome1Response301;

export type adminHome1Response = adminHome1ResponseComposite & {
  headers: Headers;
};

export const getAdminHome1Url = () => {
  return `/admin`;
};

export const adminHome1 = async (
  options?: RequestInit,
): Promise<adminHome1Response> => {
  return customInstance<adminHome1Response>(getAdminHome1Url(), {
    ...options,
    method: "GET",
  });
};

export const getAdminHome1QueryKey = () => {
  return [`/admin`] as const;
};

export const getAdminHome1InfiniteQueryOptions = <
  TData = InfiniteData<Awaited<ReturnType<typeof adminHome1>>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminHome1>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminHome1QueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminHome1>>> = ({
    signal,
  }) => adminHome1({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof adminHome1>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminHome1InfiniteQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminHome1>>
>;
export type AdminHome1InfiniteQueryError = void;

export function useAdminHome1Infinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminHome1>>>,
  TError = void,
>(options: {
  query: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminHome1>>,
      TError,
      TData
    >
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminHome1>>,
        TError,
        Awaited<ReturnType<typeof adminHome1>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminHome1Infinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminHome1>>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminHome1>>,
      TError,
      TData
    >
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminHome1>>,
        TError,
        Awaited<ReturnType<typeof adminHome1>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminHome1Infinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminHome1>>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminHome1>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary home admin
 */

export function useAdminHome1Infinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminHome1>>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminHome1>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminHome1InfiniteQueryOptions(options);

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<
    TData,
    TError
  > & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

export const getAdminHome1QueryOptions = <
  TData = Awaited<ReturnType<typeof adminHome1>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminHome1>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminHome1QueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminHome1>>> = ({
    signal,
  }) => adminHome1({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof adminHome1>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminHome1QueryResult = NonNullable<
  Awaited<ReturnType<typeof adminHome1>>
>;
export type AdminHome1QueryError = void;

export function useAdminHome1<
  TData = Awaited<ReturnType<typeof adminHome1>>,
  TError = void,
>(options: {
  query: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminHome1>>, TError, TData>
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminHome1>>,
        TError,
        Awaited<ReturnType<typeof adminHome1>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminHome1<
  TData = Awaited<ReturnType<typeof adminHome1>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminHome1>>, TError, TData>
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminHome1>>,
        TError,
        Awaited<ReturnType<typeof adminHome1>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminHome1<
  TData = Awaited<ReturnType<typeof adminHome1>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminHome1>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary home admin
 */

export function useAdminHome1<
  TData = Awaited<ReturnType<typeof adminHome1>>,
  TError = void,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminHome1>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminHome1QueryOptions(options);

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };

  query.queryKey = queryOptions.queryKey;

  return query;
}

/**
 * @summary Download ui/build
 */
export type adminAdminFilepathResponse200 = {
  data: void;
  status: 200;
};

export type adminAdminFilepathResponse404 = {
  data: void;
  status: 404;
};

export type adminAdminFilepathResponseComposite =
  | adminAdminFilepathResponse200
  | adminAdminFilepathResponse404;

export type adminAdminFilepathResponse = adminAdminFilepathResponseComposite & {
  headers: Headers;
};

export const getAdminAdminFilepathUrl = (filepath: unknown) => {
  return `/admin/${filepath}`;
};

export const adminAdminFilepath = async (
  filepath: unknown,
  options?: RequestInit,
): Promise<adminAdminFilepathResponse> => {
  return customInstance<adminAdminFilepathResponse>(
    getAdminAdminFilepathUrl(filepath),
    {
      ...options,
      method: "GET",
    },
  );
};

export const getAdminAdminFilepathQueryKey = (filepath: unknown) => {
  return [`/admin/${filepath}`] as const;
};

export const getAdminAdminFilepathInfiniteQueryOptions = <
  TData = InfiniteData<Awaited<ReturnType<typeof adminAdminFilepath>>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseInfiniteQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey =
    queryOptions?.queryKey ?? getAdminAdminFilepathQueryKey(filepath);

  const queryFn: QueryFunction<
    Awaited<ReturnType<typeof adminAdminFilepath>>
  > = ({ signal }) =>
    adminAdminFilepath(filepath, { signal, ...requestOptions });

  return {
    queryKey,
    queryFn,
    enabled: !!filepath,
    ...queryOptions,
  } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof adminAdminFilepath>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminAdminFilepathInfiniteQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminAdminFilepath>>
>;
export type AdminAdminFilepathInfiniteQueryError = void;

export function useAdminAdminFilepathInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminAdminFilepath>>>,
  TError = void,
>(
  filepath: unknown,
  options: {
    query: Partial<
      UseInfiniteQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    > &
      Pick<
        DefinedInitialDataOptions<
          Awaited<ReturnType<typeof adminAdminFilepath>>,
          TError,
          Awaited<ReturnType<typeof adminAdminFilepath>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customInstance>;
  },
): DefinedUseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminAdminFilepathInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminAdminFilepath>>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseInfiniteQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    > &
      Pick<
        UndefinedInitialDataOptions<
          Awaited<ReturnType<typeof adminAdminFilepath>>,
          TError,
          Awaited<ReturnType<typeof adminAdminFilepath>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminAdminFilepathInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminAdminFilepath>>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseInfiniteQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download ui/build
 */

export function useAdminAdminFilepathInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminAdminFilepath>>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseInfiniteQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminAdminFilepathInfiniteQueryOptions(
    filepath,
    options,
  );

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<
    TData,
    TError
  > & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

export const getAdminAdminFilepathQueryOptions = <
  TData = Awaited<ReturnType<typeof adminAdminFilepath>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey =
    queryOptions?.queryKey ?? getAdminAdminFilepathQueryKey(filepath);

  const queryFn: QueryFunction<
    Awaited<ReturnType<typeof adminAdminFilepath>>
  > = ({ signal }) =>
    adminAdminFilepath(filepath, { signal, ...requestOptions });

  return {
    queryKey,
    queryFn,
    enabled: !!filepath,
    ...queryOptions,
  } as UseQueryOptions<
    Awaited<ReturnType<typeof adminAdminFilepath>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminAdminFilepathQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminAdminFilepath>>
>;
export type AdminAdminFilepathQueryError = void;

export function useAdminAdminFilepath<
  TData = Awaited<ReturnType<typeof adminAdminFilepath>>,
  TError = void,
>(
  filepath: unknown,
  options: {
    query: Partial<
      UseQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    > &
      Pick<
        DefinedInitialDataOptions<
          Awaited<ReturnType<typeof adminAdminFilepath>>,
          TError,
          Awaited<ReturnType<typeof adminAdminFilepath>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customInstance>;
  },
): DefinedUseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminAdminFilepath<
  TData = Awaited<ReturnType<typeof adminAdminFilepath>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    > &
      Pick<
        UndefinedInitialDataOptions<
          Awaited<ReturnType<typeof adminAdminFilepath>>,
          TError,
          Awaited<ReturnType<typeof adminAdminFilepath>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminAdminFilepath<
  TData = Awaited<ReturnType<typeof adminAdminFilepath>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download ui/build
 */

export function useAdminAdminFilepath<
  TData = Awaited<ReturnType<typeof adminAdminFilepath>>,
  TError = void,
>(
  filepath: unknown,
  options?: {
    query?: Partial<
      UseQueryOptions<
        Awaited<ReturnType<typeof adminAdminFilepath>>,
        TError,
        TData
      >
    >;
    request?: SecondParameter<typeof customInstance>;
  },
): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminAdminFilepathQueryOptions(filepath, options);

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };

  query.queryKey = queryOptions.queryKey;

  return query;
}

/**
 * @summary Download admin/public/favicon.ico
 */
export type adminFaviconIcoResponse200 = {
  data: void;
  status: 200;
};

export type adminFaviconIcoResponseComposite = adminFaviconIcoResponse200;

export type adminFaviconIcoResponse = adminFaviconIcoResponseComposite & {
  headers: Headers;
};

export const getAdminFaviconIcoUrl = () => {
  return `/favicon.ico`;
};

export const adminFaviconIco = async (
  options?: RequestInit,
): Promise<adminFaviconIcoResponse> => {
  return customInstance<adminFaviconIcoResponse>(getAdminFaviconIcoUrl(), {
    ...options,
    method: "GET",
  });
};

export const getAdminFaviconIcoQueryKey = () => {
  return [`/favicon.ico`] as const;
};

export const getAdminFaviconIcoInfiniteQueryOptions = <
  TData = InfiniteData<Awaited<ReturnType<typeof adminFaviconIco>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminFaviconIco>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminFaviconIcoQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminFaviconIco>>> = ({
    signal,
  }) => adminFaviconIco({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof adminFaviconIco>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminFaviconIcoInfiniteQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminFaviconIco>>
>;
export type AdminFaviconIcoInfiniteQueryError = unknown;

export function useAdminFaviconIcoInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminFaviconIco>>>,
  TError = unknown,
>(options: {
  query: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminFaviconIco>>,
      TError,
      TData
    >
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminFaviconIco>>,
        TError,
        Awaited<ReturnType<typeof adminFaviconIco>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminFaviconIcoInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminFaviconIco>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminFaviconIco>>,
      TError,
      TData
    >
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminFaviconIco>>,
        TError,
        Awaited<ReturnType<typeof adminFaviconIco>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminFaviconIcoInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminFaviconIco>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminFaviconIco>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download admin/public/favicon.ico
 */

export function useAdminFaviconIcoInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminFaviconIco>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminFaviconIco>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminFaviconIcoInfiniteQueryOptions(options);

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<
    TData,
    TError
  > & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

export const getAdminFaviconIcoQueryOptions = <
  TData = Awaited<ReturnType<typeof adminFaviconIco>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminFaviconIco>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminFaviconIcoQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminFaviconIco>>> = ({
    signal,
  }) => adminFaviconIco({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof adminFaviconIco>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminFaviconIcoQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminFaviconIco>>
>;
export type AdminFaviconIcoQueryError = unknown;

export function useAdminFaviconIco<
  TData = Awaited<ReturnType<typeof adminFaviconIco>>,
  TError = unknown,
>(options: {
  query: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminFaviconIco>>, TError, TData>
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminFaviconIco>>,
        TError,
        Awaited<ReturnType<typeof adminFaviconIco>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminFaviconIco<
  TData = Awaited<ReturnType<typeof adminFaviconIco>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminFaviconIco>>, TError, TData>
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminFaviconIco>>,
        TError,
        Awaited<ReturnType<typeof adminFaviconIco>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminFaviconIco<
  TData = Awaited<ReturnType<typeof adminFaviconIco>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminFaviconIco>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download admin/public/favicon.ico
 */

export function useAdminFaviconIco<
  TData = Awaited<ReturnType<typeof adminFaviconIco>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminFaviconIco>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminFaviconIcoQueryOptions(options);

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };

  query.queryKey = queryOptions.queryKey;

  return query;
}

/**
 * @summary Download admin/public/robots.txt
 */
export type adminRobotsTxtResponse200 = {
  data: void;
  status: 200;
};

export type adminRobotsTxtResponseComposite = adminRobotsTxtResponse200;

export type adminRobotsTxtResponse = adminRobotsTxtResponseComposite & {
  headers: Headers;
};

export const getAdminRobotsTxtUrl = () => {
  return `/robots.txt`;
};

export const adminRobotsTxt = async (
  options?: RequestInit,
): Promise<adminRobotsTxtResponse> => {
  return customInstance<adminRobotsTxtResponse>(getAdminRobotsTxtUrl(), {
    ...options,
    method: "GET",
  });
};

export const getAdminRobotsTxtQueryKey = () => {
  return [`/robots.txt`] as const;
};

export const getAdminRobotsTxtInfiniteQueryOptions = <
  TData = InfiniteData<Awaited<ReturnType<typeof adminRobotsTxt>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminRobotsTxt>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminRobotsTxtQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminRobotsTxt>>> = ({
    signal,
  }) => adminRobotsTxt({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof adminRobotsTxt>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminRobotsTxtInfiniteQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminRobotsTxt>>
>;
export type AdminRobotsTxtInfiniteQueryError = unknown;

export function useAdminRobotsTxtInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminRobotsTxt>>>,
  TError = unknown,
>(options: {
  query: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminRobotsTxt>>,
      TError,
      TData
    >
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminRobotsTxt>>,
        TError,
        Awaited<ReturnType<typeof adminRobotsTxt>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminRobotsTxtInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminRobotsTxt>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminRobotsTxt>>,
      TError,
      TData
    >
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminRobotsTxt>>,
        TError,
        Awaited<ReturnType<typeof adminRobotsTxt>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminRobotsTxtInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminRobotsTxt>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminRobotsTxt>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download admin/public/robots.txt
 */

export function useAdminRobotsTxtInfinite<
  TData = InfiniteData<Awaited<ReturnType<typeof adminRobotsTxt>>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseInfiniteQueryOptions<
      Awaited<ReturnType<typeof adminRobotsTxt>>,
      TError,
      TData
    >
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseInfiniteQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminRobotsTxtInfiniteQueryOptions(options);

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<
    TData,
    TError
  > & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

export const getAdminRobotsTxtQueryOptions = <
  TData = Awaited<ReturnType<typeof adminRobotsTxt>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminRobotsTxt>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getAdminRobotsTxtQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminRobotsTxt>>> = ({
    signal,
  }) => adminRobotsTxt({ signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof adminRobotsTxt>>,
    TError,
    TData
  > & { queryKey: DataTag<QueryKey, TData, TError> };
};

export type AdminRobotsTxtQueryResult = NonNullable<
  Awaited<ReturnType<typeof adminRobotsTxt>>
>;
export type AdminRobotsTxtQueryError = unknown;

export function useAdminRobotsTxt<
  TData = Awaited<ReturnType<typeof adminRobotsTxt>>,
  TError = unknown,
>(options: {
  query: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminRobotsTxt>>, TError, TData>
  > &
    Pick<
      DefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminRobotsTxt>>,
        TError,
        Awaited<ReturnType<typeof adminRobotsTxt>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): DefinedUseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminRobotsTxt<
  TData = Awaited<ReturnType<typeof adminRobotsTxt>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminRobotsTxt>>, TError, TData>
  > &
    Pick<
      UndefinedInitialDataOptions<
        Awaited<ReturnType<typeof adminRobotsTxt>>,
        TError,
        Awaited<ReturnType<typeof adminRobotsTxt>>
      >,
      "initialData"
    >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
export function useAdminRobotsTxt<
  TData = Awaited<ReturnType<typeof adminRobotsTxt>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminRobotsTxt>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
};
/**
 * @summary Download admin/public/robots.txt
 */

export function useAdminRobotsTxt<
  TData = Awaited<ReturnType<typeof adminRobotsTxt>>,
  TError = unknown,
>(options?: {
  query?: Partial<
    UseQueryOptions<Awaited<ReturnType<typeof adminRobotsTxt>>, TError, TData>
  >;
  request?: SecondParameter<typeof customInstance>;
}): UseQueryResult<TData, TError> & {
  queryKey: DataTag<QueryKey, TData, TError>;
} {
  const queryOptions = getAdminRobotsTxtQueryOptions(options);

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };

  query.queryKey = queryOptions.queryKey;

  return query;
}
