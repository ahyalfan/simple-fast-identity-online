-- Active: 1723868453146@@127.0.0.1@5432@fido
CREATE DATABASE fido;

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;

SET lock_timeout = 0;

SET idle_in_transaction_session_timeout = 0;

SET client_encoding = 'UTF8';

SET standard_conforming_strings = on;

SELECT pg_catalog.set_config ('search_path', '', false);

SET check_function_bodies = false;

SET xmloption = content;

SET client_min_messages = warning;

SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: challenges; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.challenge (
    id character varying(100) NOT NULL,
    key character varying(255) NOT NULL,
    expired_at bigint NOT NULL,
    validated_at bigint DEFAULT 0 NOT NULL
);

--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    public_key character varying(255) NOT NULL,
    created_at bigint NOT NULL,
    device_id character varying(100) NOT NULL
);

--
-- Data for Name: challenges; Type: TABLE DATA; Schema: public; Owner: -
--

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

--
-- Name: challenges challenges_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.challenge
ADD CONSTRAINT challenges_pk PRIMARY KEY (id);

--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
ADD CONSTRAINT users_pk PRIMARY KEY (id);

--
-- PostgreSQL database dump complete
--

SELECT * FROM users

SELECT * FROM challenge