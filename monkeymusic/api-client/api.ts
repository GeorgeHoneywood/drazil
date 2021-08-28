/* tslint:disable */
/* eslint-disable */
/**
 * spec/monkey.proto
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
    imageUrl?: string;
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
 * @interface SpecSong
 */
export interface SpecSong {
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
}
/**
 * 
 * @export
 * @interface SpecSongsReply
 */
export interface SpecSongsReply {
    /**
     * 
     * @type {Array<SpecSong>}
     * @memberof SpecSongsReply
     */
    songs?: Array<SpecSong>;
}

/**
 * MonkeyApi - axios parameter creator
 * @export
 */
export const MonkeyApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        monkeyListAlbums: async (artistId: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'artistId' is not null or undefined
            assertParamExists('monkeyListAlbums', 'artistId', artistId)
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
        monkeyListArtists: async (options: any = {}): Promise<RequestArgs> => {
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
        monkeyListSongs: async (artistId: string, albumId: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'artistId' is not null or undefined
            assertParamExists('monkeyListSongs', 'artistId', artistId)
            // verify required parameter 'albumId' is not null or undefined
            assertParamExists('monkeyListSongs', 'albumId', albumId)
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
    }
};

/**
 * MonkeyApi - functional programming interface
 * @export
 */
export const MonkeyApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = MonkeyApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async monkeyListAlbums(artistId: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecAlbumsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.monkeyListAlbums(artistId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async monkeyListArtists(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecArtistsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.monkeyListArtists(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} artistId 
         * @param {string} albumId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async monkeyListSongs(artistId: string, albumId: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SpecSongsReply>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.monkeyListSongs(artistId, albumId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * MonkeyApi - factory interface
 * @export
 */
export const MonkeyApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = MonkeyApiFp(configuration)
    return {
        /**
         * 
         * @param {string} artistId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        monkeyListAlbums(artistId: string, options?: any): AxiosPromise<SpecAlbumsReply> {
            return localVarFp.monkeyListAlbums(artistId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        monkeyListArtists(options?: any): AxiosPromise<SpecArtistsReply> {
            return localVarFp.monkeyListArtists(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} artistId 
         * @param {string} albumId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        monkeyListSongs(artistId: string, albumId: string, options?: any): AxiosPromise<SpecSongsReply> {
            return localVarFp.monkeyListSongs(artistId, albumId, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * MonkeyApi - object-oriented interface
 * @export
 * @class MonkeyApi
 * @extends {BaseAPI}
 */
export class MonkeyApi extends BaseAPI {
    /**
     * 
     * @param {string} artistId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof MonkeyApi
     */
    public monkeyListAlbums(artistId: string, options?: any) {
        return MonkeyApiFp(this.configuration).monkeyListAlbums(artistId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof MonkeyApi
     */
    public monkeyListArtists(options?: any) {
        return MonkeyApiFp(this.configuration).monkeyListArtists(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} artistId 
     * @param {string} albumId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof MonkeyApi
     */
    public monkeyListSongs(artistId: string, albumId: string, options?: any) {
        return MonkeyApiFp(this.configuration).monkeyListSongs(artistId, albumId, options).then((request) => request(this.axios, this.basePath));
    }
}


