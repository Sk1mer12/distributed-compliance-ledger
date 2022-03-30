/* eslint-disable */
import { Certificate } from '../pki/certificate';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseRevokedCertificates = { subject: '', subjectKeyId: '', subjectAsText: '' };
export const RevokedCertificates = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        for (const v of message.certs) {
            Certificate.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.subjectAsText !== '') {
            writer.uint32(34).string(message.subjectAsText);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRevokedCertificates };
        message.certs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.subject = reader.string();
                    break;
                case 2:
                    message.subjectKeyId = reader.string();
                    break;
                case 3:
                    message.certs.push(Certificate.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.subjectAsText = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRevokedCertificates };
        message.certs = [];
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.certs !== undefined && object.certs !== null) {
            for (const e of object.certs) {
                message.certs.push(Certificate.fromJSON(e));
            }
        }
        if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
            message.subjectAsText = String(object.subjectAsText);
        }
        else {
            message.subjectAsText = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        if (message.certs) {
            obj.certs = message.certs.map((e) => (e ? Certificate.toJSON(e) : undefined));
        }
        else {
            obj.certs = [];
        }
        message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRevokedCertificates };
        message.certs = [];
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.certs !== undefined && object.certs !== null) {
            for (const e of object.certs) {
                message.certs.push(Certificate.fromPartial(e));
            }
        }
        if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
            message.subjectAsText = object.subjectAsText;
        }
        else {
            message.subjectAsText = '';
        }
        return message;
    }
};
