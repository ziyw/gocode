--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 14.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: note; Type: TABLE; Schema: public; Owner: ziyan
--

CREATE TABLE public.note (
    id integer NOT NULL,
    title character varying NOT NULL,
    content character varying NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.note OWNER TO ziyan;

--
-- Name: note_id_seq; Type: SEQUENCE; Schema: public; Owner: ziyan
--

CREATE SEQUENCE public.note_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.note_id_seq OWNER TO ziyan;

--
-- Name: note_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ziyan
--

ALTER SEQUENCE public.note_id_seq OWNED BY public.note.id;


--
-- Name: note id; Type: DEFAULT; Schema: public; Owner: ziyan
--

ALTER TABLE ONLY public.note ALTER COLUMN id SET DEFAULT nextval('public.note_id_seq'::regclass);


--
-- Data for Name: note; Type: TABLE DATA; Schema: public; Owner: ziyan
--

COPY public.note (id, title, content, created_at) FROM stdin;
1	First Entry	The very first note entry	2023-09-13 20:50:00.105541
2	Entry 2	Second try	2023-09-13 20:50:11.487881
3	Entry three	Third time is a charm	2023-09-13 20:50:57.312694
\.


--
-- Name: note_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ziyan
--

SELECT pg_catalog.setval('public.note_id_seq', 3, true);


--
-- Name: note note_pkey; Type: CONSTRAINT; Schema: public; Owner: ziyan
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT note_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

