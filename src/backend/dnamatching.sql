--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.2

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

--
-- Name: deuq9ceh711el7; Type: DATABASE; Schema: -; Owner: azvieqzbjnzkmn
--

CREATE DATABASE deuq9ceh711el7 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE deuq9ceh711el7 OWNER TO azvieqzbjnzkmn;

\connect deuq9ceh711el7

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
-- Name: disease; Type: TABLE; Schema: public; Owner: azvieqzbjnzkmn
--

CREATE TABLE public.disease (
    name character varying(50) NOT NULL,
    dna character varying(100) DEFAULT NULL::character varying
);


ALTER TABLE public.disease OWNER TO azvieqzbjnzkmn;

--
-- Name: history; Type: TABLE; Schema: public; Owner: azvieqzbjnzkmn
--

CREATE TABLE public.history (
    id integer NOT NULL,
    date character varying(10) DEFAULT NULL::character varying,
    name character varying(50) DEFAULT NULL::character varying,
    disease character varying(50) DEFAULT NULL::character varying,
    result smallint,
    similarity integer NOT NULL
);


ALTER TABLE public.history OWNER TO azvieqzbjnzkmn;

--
-- Name: history_id_seq; Type: SEQUENCE; Schema: public; Owner: azvieqzbjnzkmn
--

CREATE SEQUENCE public.history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.history_id_seq OWNER TO azvieqzbjnzkmn;

--
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: azvieqzbjnzkmn
--

ALTER SEQUENCE public.history_id_seq OWNED BY public.history.id;


--
-- Name: history id; Type: DEFAULT; Schema: public; Owner: azvieqzbjnzkmn
--

ALTER TABLE ONLY public.history ALTER COLUMN id SET DEFAULT nextval('public.history_id_seq'::regclass);


--
-- Data for Name: disease; Type: TABLE DATA; Schema: public; Owner: azvieqzbjnzkmn
--

COPY public.disease (name, dna) FROM stdin;
cacarair	ATGGTGCACGAT
flu	ATGCTGACGAT
COVID-19	CACATAGATTG
thalasemia	AAAAA
turner	TGGCACTG
klinefelter	CCGGCGTACGCGTCCCATAT
diabetes	AAACCTGTCATAACTTACCT
sickle cell	GAGACTACTTGGAAATGTGG
down syndrome	CTAGATCTTTGCCCACGCAC
hemofilia	CTAATCGGTCCACGTTTGGT
kanker payudara	GCGGGTACTAGATGA
kanker serviks	CTGCAGGGACTCCGA
alzheimer	CGTTAAGTACATTAC
autis	CCCGTCATAGGCGCC
migrain	GTTCAGGATCACGTT
artritis	ACCGCCATAAGATGG
trisomy 18	GAGCATGACTTCTTC
trisomy 13	TCCGCTGCGCCCACG
ambis	CCTATAACCCTTCTG
bucin	TCCGCTGCGCCCACG
Malaria	TTTTTATGGAGCTCG
Klinefelter	GCATGATATTTAGACGATGATG
\.


--
-- Data for Name: history; Type: TABLE DATA; Schema: public; Owner: azvieqzbjnzkmn
--

COPY public.history (id, date, name, disease, result, similarity) FROM stdin;
1	2022-03-01	saul	flu	0	0
2	2022-01-30	afan	cacarair	0	0
3	2022-04-28	Ahmad Alfani Handoyo	COVID-19	1	81
4	2022-04-28	Jova	COVID-19	1	100
5	2022-04-28	Jova	COVID-19	1	100
6	2022-04-28	Saul	kanker payudara	1	100
7	2022-04-28	Saul	kanker payudara	1	100
8	2022-04-28	Liza	autis	0	73
\.


--
-- Name: history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: azvieqzbjnzkmn
--

SELECT pg_catalog.setval('public.history_id_seq', 8, true);


--
-- Name: disease disease_pkey; Type: CONSTRAINT; Schema: public; Owner: azvieqzbjnzkmn
--

ALTER TABLE ONLY public.disease
    ADD CONSTRAINT disease_pkey PRIMARY KEY (name);


--
-- Name: history history_pkey; Type: CONSTRAINT; Schema: public; Owner: azvieqzbjnzkmn
--

ALTER TABLE ONLY public.history
    ADD CONSTRAINT history_pkey PRIMARY KEY (id);


--
-- Name: history history_ibfk_1; Type: FK CONSTRAINT; Schema: public; Owner: azvieqzbjnzkmn
--

ALTER TABLE ONLY public.history
    ADD CONSTRAINT history_ibfk_1 FOREIGN KEY (disease) REFERENCES public.disease(name);


--
-- Name: DATABASE deuq9ceh711el7; Type: ACL; Schema: -; Owner: azvieqzbjnzkmn
--

REVOKE CONNECT,TEMPORARY ON DATABASE deuq9ceh711el7 FROM PUBLIC;


--
-- Name: LANGUAGE plpgsql; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON LANGUAGE plpgsql TO azvieqzbjnzkmn;


--
-- PostgreSQL database dump complete
--

