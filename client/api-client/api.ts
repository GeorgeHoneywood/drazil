/* tslint:disable */
/* eslint-disable */
/**
 * spec/drazil.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface ProtobufAny
 */
export interface ProtobufAny {
    /**
     * 
     * @type {string}
     * @memberof ProtobufAny
     */
    typeUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ProtobufAny
     */
    value?: string;
}
/**
 * 
 * @export
 * @interface RpcStatus
 */
export interface RpcStatus {
    /**
     * 
     * @type {number}
     * @memberof RpcStatus
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof RpcStatus
     */
    message?: string;
    /**
     * 
     * @type {Array<ProtobufAny>}
     * @memberof RpcStatus
     */
    details?: Array<ProtobufAny>;
}
/**
 * 
 * @export
 * @interface SpecAlbum
 */
export interface SpecAlbum {
    /**
     * 
     * @type {string}
     * @memberof SpecAlbum
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAlbum
     */
    name?: string;
}
/**
 * 
 * @export
 * @interface SpecAlbumsReply
 */
export interface SpecAlbumsReply {
    /**
     * 
     * @type {Array<SpecAlbum>}
     * @memberof SpecAlbumsReply
     */
    albums?: Array<SpecAlbum>;
    /**
     * 
     * @type {string}
     * @memberof SpecAlbumsReply
     */
    artistName?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAlbumsReply
     */
    artistId?: string;
}
/**
 * 
 * @export
 * @interface SpecAllAlbum
 */
export interface SpecAllAlbum {
    /**
     * 
     * @type {string}
     * @memberof SpecAllAlbum
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllAlbum
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllAlbum
     */
    artistName?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllAlbum
     */
    artistId?: string;
}
/**
 * 
 * @export
 * @interface SpecAllAlbumsReply
 */
export interface SpecAllAlbumsReply {
    /**
     * 
     * @type {Array<SpecAllAlbum>}
     * @memberof SpecAllAlbumsReply
     */
    albums?: Array<SpecAllAlbum>;
}
/**
 * 
 * @export
 * @interface SpecAllSong
 */
export interface SpecAllSong {
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    albumId?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    albumName?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    artistName?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    artistId?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof SpecAllSong
     */
    number?: number;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    lyrics?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecAllSong
     */
    fileType?: string;
    /**
     * 
     * @type {number}
     * @memberof SpecAllSong
     */
    year?: number;
}
/**
 * 
 * @export
 * @interface SpecArtist
 */
export interface SpecArtist {
    /**
     * 
     * @type {string}
     * @memberof SpecArtist
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecArtist
     */
    imageUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecArtist
     */
    id?: string;
}
/**
 * 
 * @export
 * @interface SpecArtistsReply
 */
export interface SpecArtistsReply {
    /**
     * 
     * @type {Array<SpecArtist>}
     * @memberof SpecArtistsReply
     */
    artists?: Array<SpecArtist>;
    /**
     * 
     * @type {number}
     * @memberof SpecArtistsReply
     */
    artistCount?: number;
}
/**
 * 
 * @export
 * @interface SpecSearchReply
 */
export interface SpecSearchReply {
    /**
     * 
     * @type {Array<SpecArtist>}
     * @memberof SpecSearchReply
     */
    artists?: Array<SpecArtist>;
    /**
     * 
     * @type {Array<SpecAllAlbum>}
     * @memberof SpecSearchReply
     */
    albums?: Array<SpecAllAlbum>;
    /**
     * 
     * @type {Array<SpecAllSong>}
     * @memberof SpecSearchReply
     */
    songs?: Array<SpecAllSong>;
}
/**
 * 
 * @export
 * @interface SpecSong
 */
export interface SpecSong {
    /**
     * 
     * @type {string}
     * @memberof SpecSong
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecSong
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof SpecSong
     */
    number?: number;
    /**
     * 
     * @type {string}
     * @memberof SpecSong
     */
    lyrics?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecSong
     */
    fileType?: string;
    /**
     * 
     * @type {number}
     * @memberof SpecSong
     */
    year?: number;
}
/**
 * 
 * @export
 * @interface SpecSongsReply
 */
export interface SpecSongsReply {
    /**
     * 
     * @type {string}
     * @memberof SpecSongsReply
     */
    artistName?: string;
    /**
     * 
     * @type {string}
     * @memberof SpecSongsReply
     */
    albumName?: string;
    /**
     * 
     * @type {Array<SpecSong>}
     * @memberof SpecSongsReply
     */
    songs?: Array<SpecSong>;
}

/**
 * DrazilApi - axios parameter creator
 * @export
 */
export const DrazilApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListAlbums: async (artistId: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'artistId' is not null or undefined
            assertParamExists('drazilListAlbums', 'artistId', artistId)
            const localVarPath = `/v1/artist/{artistId}/albums`
                .replace(`{${"artistId"}}`, encodeURIComponent(String(artistId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListAllAlbums: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/v1/albums`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListArtists: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/v1/artists`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} artistId 
         * @param {string} albumId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListSongs: async (artistId: string, albumId: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'artistId' is not null or undefined
            assertParamExists('drazilListSongs', 'artistId', artistId)
            // verify required parameter 'albumId' is not null or undefined
            assertParamExists('drazilListSongs', 'albumId', albumId)
            const localVarPath = `/v1/artist/{artistId}/album/{albumId}/songs`
                .replace(`{${"artistId"}}`, encodeURIComponent(String(artistId)))
                .replace(`{${"albumId"}}`, encodeURIComponent(String(albumId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} token 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilSearch: async (token: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'token' is not null or undefined
            assertParamExists('drazilSearch', 'token', token)
            const localVarPath = `/v1/search/{token}`
                .replace(`{${"token"}}`, encodeURIComponent(String(token)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DrazilApi - functional programming interface
 * @export
 */
export const DrazilApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DrazilApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async drazilListAlbums(artistId: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecAlbumsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.drazilListAlbums(artistId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async drazilListAllAlbums(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecAllAlbumsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.drazilListAllAlbums(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async drazilListArtists(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecArtistsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.drazilListArtists(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} artistId 
         * @param {string} albumId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async drazilListSongs(artistId: string, albumId: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecSongsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.drazilListSongs(artistId, albumId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} token 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async drazilSearch(token: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecSearchReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.drazilSearch(token, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DrazilApi - factory interface
 * @export
 */
export const DrazilApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DrazilApiFp(configuration)
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListAlbums(artistId: string, options?: any): AxiosPromise<SpecAlbumsReply> {
            return localVarFp.drazilListAlbums(artistId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListAllAlbums(options?: any): AxiosPromise<SpecAllAlbumsReply> {
            return localVarFp.drazilListAllAlbums(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListArtists(options?: any): AxiosPromise<SpecArtistsReply> {
            return localVarFp.drazilListArtists(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} artistId 
         * @param {string} albumId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilListSongs(artistId: string, albumId: string, options?: any): AxiosPromise<SpecSongsReply> {
            return localVarFp.drazilListSongs(artistId, albumId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} token 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        drazilSearch(token: string, options?: any): AxiosPromise<SpecSearchReply> {
            return localVarFp.drazilSearch(token, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DrazilApi - object-oriented interface
 * @export
 * @class DrazilApi
 * @extends {BaseAPI}
 */
export class DrazilApi extends BaseAPI {
    /**
     * 
     * @param {string} artistId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DrazilApi
     */
    public drazilListAlbums(artistId: string, options?: any) {
        return DrazilApiFp(this.configuration).drazilListAlbums(artistId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DrazilApi
     */
    public drazilListAllAlbums(options?: any) {
        return DrazilApiFp(this.configuration).drazilListAllAlbums(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DrazilApi
     */
    public drazilListArtists(options?: any) {
        return DrazilApiFp(this.configuration).drazilListArtists(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} artistId 
     * @param {string} albumId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DrazilApi
     */
    public drazilListSongs(artistId: string, albumId: string, options?: any) {
        return DrazilApiFp(this.configuration).drazilListSongs(artistId, albumId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} token 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DrazilApi
     */
    public drazilSearch(token: string, options?: any) {
        return DrazilApiFp(this.configuration).drazilSearch(token, options).then((request) => request(this.axios, this.basePath));
    }
}

