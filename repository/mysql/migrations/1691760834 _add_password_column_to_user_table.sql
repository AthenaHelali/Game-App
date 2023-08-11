-- +migrate Up
ALTER TABLE users add column password text;