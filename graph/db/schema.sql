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
-- Name: edge; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.edge (
    id integer NOT NULL,
    src integer NOT NULL,
    dst integer NOT NULL
);


--
-- Name: edge_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.edge ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.edge_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: node; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.node (
    id integer NOT NULL,
    label character varying NOT NULL
);


--
-- Name: node_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.node ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.node_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: edge edge_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.edge
    ADD CONSTRAINT edge_pkey PRIMARY KEY (id);


--
-- Name: node node_label_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_label_key UNIQUE (label);


--
-- Name: node node_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: edge fk_node_dst; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.edge
    ADD CONSTRAINT fk_node_dst FOREIGN KEY (dst) REFERENCES public.node(id);


--
-- Name: edge fk_node_src; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.edge
    ADD CONSTRAINT fk_node_src FOREIGN KEY (src) REFERENCES public.node(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240923203420');
