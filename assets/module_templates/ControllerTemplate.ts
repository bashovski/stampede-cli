import Controller from './Controller.ts';
import __replace_me__Service from '../services/__replace_me__Service.ts';
import HttpResponse from "../http/HttpResponse.ts";
import Logger from '../lib/Logger.ts';

/**
 * @class __replace_me__Controller
 * @summary Handles all __replace_me__-related HTTP requests
 */
class __replace_me__Controller extends Controller {

    /**
     * @summary Handles index request
     * @param ctx
     */
    static async index(ctx: any): Promise<void> {
        try {

            const { response, error } : any = await __replace_me__Service.index();
            (response || error).send(ctx.response);

        } catch(error) {
            Logger.error(error);
            this.sendDefaultError(ctx);
        }
    }

    /**
     * @summary Sends default (400 - Bad request) HTTP error to the client
     * @param ctx
     */
    private static sendDefaultError(ctx: any): void {
        new HttpResponse(400, {
            error: 'Bad request'
        }).send(ctx.response);
    }
}

/**
 * @exports __replace_me__
 */
export default __replace_me__Controller;
