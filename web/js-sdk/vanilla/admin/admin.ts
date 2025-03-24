/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import { customInstance } from "../../mutator/custom-instance";

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
