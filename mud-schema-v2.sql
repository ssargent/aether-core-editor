--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Postgres.app)
-- Dumped by pg_dump version 17.0

-- Started on 2025-05-28 18:47:47 EDT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- TOC entry 223 (class 1259 OID 19125)
-- Name: attributes; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.attributes (
    id bigint NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.attributes OWNER TO muddy;

--
-- TOC entry 222 (class 1259 OID 19124)
-- Name: attributes_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.attributes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.attributes_id_seq OWNER TO muddy;

--
-- TOC entry 3847 (class 0 OID 0)
-- Dependencies: 222
-- Name: attributes_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.attributes_id_seq OWNED BY game.attributes.id;


--
-- TOC entry 264 (class 1259 OID 19663)
-- Name: capabilities; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.capabilities (
    id bigint NOT NULL,
    world_id bigint NOT NULL,
    parent_id bigint,
    capability_type text NOT NULL,
    code text NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    requirements jsonb DEFAULT '{}'::jsonb NOT NULL,
    actions jsonb,
    access_requirements jsonb DEFAULT '{}'::jsonb NOT NULL,
    tags text[] DEFAULT '{}'::text[] NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    gameplay_definition jsonb,
    CONSTRAINT capabilities_capability_type_check CHECK ((capability_type = ANY (ARRAY['skill'::text, 'feat'::text, 'ability'::text, 'power'::text, 'action'::text, 'feature'::text])))
);


ALTER TABLE game.capabilities OWNER TO muddy;

--
-- TOC entry 263 (class 1259 OID 19662)
-- Name: capabilities_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.capabilities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.capabilities_id_seq OWNER TO muddy;

--
-- TOC entry 3848 (class 0 OID 0)
-- Dependencies: 263
-- Name: capabilities_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.capabilities_id_seq OWNED BY game.capabilities.id;


--
-- TOC entry 250 (class 1259 OID 19495)
-- Name: character_class_features; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.character_class_features (
    id bigint NOT NULL,
    class_id bigint NOT NULL,
    level integer NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    code character varying(32) NOT NULL
);


ALTER TABLE game.character_class_features OWNER TO muddy;

--
-- TOC entry 249 (class 1259 OID 19494)
-- Name: character_class_features_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.character_class_features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.character_class_features_id_seq OWNER TO muddy;

--
-- TOC entry 3849 (class 0 OID 0)
-- Dependencies: 249
-- Name: character_class_features_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.character_class_features_id_seq OWNED BY game.character_class_features.id;


--
-- TOC entry 248 (class 1259 OID 19477)
-- Name: character_classes; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.character_classes (
    id bigint NOT NULL,
    world_id bigint NOT NULL,
    code character varying(32) NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    hit_points integer NOT NULL,
    stamina_expression text NOT NULL,
    skillpoint_expression text NOT NULL,
    proficiencies jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.character_classes OWNER TO muddy;

--
-- TOC entry 247 (class 1259 OID 19476)
-- Name: character_classes_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.character_classes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.character_classes_id_seq OWNER TO muddy;

--
-- TOC entry 3850 (class 0 OID 0)
-- Dependencies: 247
-- Name: character_classes_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.character_classes_id_seq OWNED BY game.character_classes.id;


--
-- TOC entry 244 (class 1259 OID 19437)
-- Name: currency; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.currency (
    id bigint NOT NULL,
    world_id bigint NOT NULL,
    code character varying(32) NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    is_spendable boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.currency OWNER TO muddy;

--
-- TOC entry 239 (class 1259 OID 19373)
-- Name: enemies; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.enemies (
    id bigint NOT NULL,
    world_id bigint NOT NULL,
    code character varying(32) NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    class character varying(32) NOT NULL,
    level integer NOT NULL,
    hit_points integer NOT NULL,
    stamina integer NOT NULL,
    strength integer NOT NULL,
    dexterity integer NOT NULL,
    constitution integer NOT NULL,
    intelligence integer NOT NULL,
    wisdom integer NOT NULL,
    weapons jsonb NOT NULL,
    armor jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.enemies OWNER TO muddy;

--
-- TOC entry 238 (class 1259 OID 19372)
-- Name: enemies_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.enemies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.enemies_id_seq OWNER TO muddy;

--
-- TOC entry 3851 (class 0 OID 0)
-- Dependencies: 238
-- Name: enemies_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.enemies_id_seq OWNED BY game.enemies.id;


--
-- TOC entry 232 (class 1259 OID 19203)
-- Name: item_categories; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.item_categories (
    id bigint NOT NULL,
    parent_id bigint,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.item_categories OWNER TO muddy;

--
-- TOC entry 231 (class 1259 OID 19202)
-- Name: item_categories_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.item_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.item_categories_id_seq OWNER TO muddy;

--
-- TOC entry 3852 (class 0 OID 0)
-- Dependencies: 231
-- Name: item_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.item_categories_id_seq OWNED BY game.item_categories.id;


--
-- TOC entry 234 (class 1259 OID 19219)
-- Name: items; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.items (
    id bigint NOT NULL,
    category_id bigint NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    item_properties jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    item_type character varying(32) NOT NULL,
    code character varying(32) NOT NULL,
    base_price bigint DEFAULT 0 NOT NULL,
    world_id bigint DEFAULT 1 NOT NULL
);


ALTER TABLE game.items OWNER TO muddy;

--
-- TOC entry 233 (class 1259 OID 19218)
-- Name: items_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.items_id_seq OWNER TO muddy;

--
-- TOC entry 3853 (class 0 OID 0)
-- Dependencies: 233
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.items_id_seq OWNED BY game.items.id;


--
-- TOC entry 230 (class 1259 OID 19182)
-- Name: npc_spawn_rules; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.npc_spawn_rules (
    npc_template_id bigint NOT NULL,
    world_node_id bigint NOT NULL,
    spawn_chance integer NOT NULL,
    spawn_quantity_min integer DEFAULT 1 NOT NULL,
    spawn_quantity_max integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE game.npc_spawn_rules OWNER TO muddy;

--
-- TOC entry 262 (class 1259 OID 19653)
-- Name: npc_spawn_rules_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.npc_spawn_rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.npc_spawn_rules_id_seq OWNER TO muddy;

--
-- TOC entry 3854 (class 0 OID 0)
-- Dependencies: 262
-- Name: npc_spawn_rules_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.npc_spawn_rules_id_seq OWNED BY game.npc_spawn_rules.id;


--
-- TOC entry 229 (class 1259 OID 19172)
-- Name: npc_templates; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.npc_templates (
    id bigint NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    npc_properties jsonb NOT NULL,
    can_spawn_multiple boolean NOT NULL,
    can_respawn boolean NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.npc_templates OWNER TO muddy;

--
-- TOC entry 228 (class 1259 OID 19171)
-- Name: npc_templates_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.npc_templates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.npc_templates_id_seq OWNER TO muddy;

--
-- TOC entry 3855 (class 0 OID 0)
-- Dependencies: 228
-- Name: npc_templates_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.npc_templates_id_seq OWNED BY game.npc_templates.id;


--
-- TOC entry 237 (class 1259 OID 19307)
-- Name: races; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.races (
    id bigint NOT NULL,
    world_id bigint NOT NULL,
    code character varying(32) NOT NULL,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE game.races OWNER TO muddy;

--
-- TOC entry 236 (class 1259 OID 19306)
-- Name: races_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.races_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.races_id_seq OWNER TO muddy;

--
-- TOC entry 3856 (class 0 OID 0)
-- Dependencies: 236
-- Name: races_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.races_id_seq OWNED BY game.races.id;


--
-- TOC entry 227 (class 1259 OID 19156)
-- Name: world_node_features; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.world_node_features (
    world_node_id bigint NOT NULL,
    feature_name character varying(32) NOT NULL,
    feature_value text NOT NULL,
    feature_properties jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE game.world_node_features OWNER TO muddy;

--
-- TOC entry 261 (class 1259 OID 19643)
-- Name: world_node_features_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.world_node_features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.world_node_features_id_seq OWNER TO muddy;

--
-- TOC entry 3857 (class 0 OID 0)
-- Dependencies: 261
-- Name: world_node_features_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.world_node_features_id_seq OWNED BY game.world_node_features.id;


--
-- TOC entry 226 (class 1259 OID 19146)
-- Name: world_nodes; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.world_nodes (
    id bigint NOT NULL,
    parent_id bigint,
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    world_id bigint NOT NULL
);


ALTER TABLE game.world_nodes OWNER TO muddy;

--
-- TOC entry 225 (class 1259 OID 19145)
-- Name: world_nodes_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.world_nodes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.world_nodes_id_seq OWNER TO muddy;

--
-- TOC entry 3858 (class 0 OID 0)
-- Dependencies: 225
-- Name: world_nodes_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.world_nodes_id_seq OWNED BY game.world_nodes.id;


--
-- TOC entry 224 (class 1259 OID 19135)
-- Name: worlds; Type: TABLE; Schema: game; Owner: muddy
--

CREATE TABLE game.worlds (
    name character varying(32) NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    id bigint NOT NULL,
    code character varying(32) NOT NULL
);


ALTER TABLE game.worlds OWNER TO muddy;

--
-- TOC entry 235 (class 1259 OID 19255)
-- Name: worlds_id_seq; Type: SEQUENCE; Schema: game; Owner: muddy
--

CREATE SEQUENCE game.worlds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE game.worlds_id_seq OWNER TO muddy;

--
-- TOC entry 3859 (class 0 OID 0)
-- Dependencies: 235
-- Name: worlds_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: muddy
--

ALTER SEQUENCE game.worlds_id_seq OWNED BY game.worlds.id;


--
-- TOC entry 3582 (class 2604 OID 19128)
-- Name: attributes id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.attributes ALTER COLUMN id SET DEFAULT nextval('game.attributes_id_seq'::regclass);


--
-- TOC entry 3625 (class 2604 OID 19666)
-- Name: capabilities id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.capabilities ALTER COLUMN id SET DEFAULT nextval('game.capabilities_id_seq'::regclass);


--
-- TOC entry 3622 (class 2604 OID 19498)
-- Name: character_class_features id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_class_features ALTER COLUMN id SET DEFAULT nextval('game.character_class_features_id_seq'::regclass);


--
-- TOC entry 3619 (class 2604 OID 19480)
-- Name: character_classes id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_classes ALTER COLUMN id SET DEFAULT nextval('game.character_classes_id_seq'::regclass);


--
-- TOC entry 3613 (class 2604 OID 19376)
-- Name: enemies id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.enemies ALTER COLUMN id SET DEFAULT nextval('game.enemies_id_seq'::regclass);


--
-- TOC entry 3602 (class 2604 OID 19206)
-- Name: item_categories id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.item_categories ALTER COLUMN id SET DEFAULT nextval('game.item_categories_id_seq'::regclass);


--
-- TOC entry 3605 (class 2604 OID 19222)
-- Name: items id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.items ALTER COLUMN id SET DEFAULT nextval('game.items_id_seq'::regclass);


--
-- TOC entry 3601 (class 2604 OID 19654)
-- Name: npc_spawn_rules id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_spawn_rules ALTER COLUMN id SET DEFAULT nextval('game.npc_spawn_rules_id_seq'::regclass);


--
-- TOC entry 3594 (class 2604 OID 19175)
-- Name: npc_templates id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_templates ALTER COLUMN id SET DEFAULT nextval('game.npc_templates_id_seq'::regclass);


--
-- TOC entry 3610 (class 2604 OID 19310)
-- Name: races id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.races ALTER COLUMN id SET DEFAULT nextval('game.races_id_seq'::regclass);


--
-- TOC entry 3593 (class 2604 OID 19644)
-- Name: world_node_features id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_node_features ALTER COLUMN id SET DEFAULT nextval('game.world_node_features_id_seq'::regclass);


--
-- TOC entry 3588 (class 2604 OID 19149)
-- Name: world_nodes id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_nodes ALTER COLUMN id SET DEFAULT nextval('game.world_nodes_id_seq'::regclass);


--
-- TOC entry 3587 (class 2604 OID 19256)
-- Name: worlds id; Type: DEFAULT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.worlds ALTER COLUMN id SET DEFAULT nextval('game.worlds_id_seq'::regclass);


--
-- TOC entry 3633 (class 2606 OID 19134)
-- Name: attributes pk_attributes_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.attributes
    ADD CONSTRAINT pk_attributes_id PRIMARY KEY (id);


--
-- TOC entry 3683 (class 2606 OID 19676)
-- Name: capabilities pk_capabilities_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.capabilities
    ADD CONSTRAINT pk_capabilities_id PRIMARY KEY (id);


--
-- TOC entry 3675 (class 2606 OID 19504)
-- Name: character_class_features pk_character_class_features_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_class_features
    ADD CONSTRAINT pk_character_class_features_id PRIMARY KEY (id);


--
-- TOC entry 3672 (class 2606 OID 19486)
-- Name: character_classes pk_character_classes_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_classes
    ADD CONSTRAINT pk_character_classes_id PRIMARY KEY (id);


--
-- TOC entry 3668 (class 2606 OID 19446)
-- Name: currency pk_currency_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.currency
    ADD CONSTRAINT pk_currency_id PRIMARY KEY (id);


--
-- TOC entry 3664 (class 2606 OID 19382)
-- Name: enemies pk_enemies_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.enemies
    ADD CONSTRAINT pk_enemies_id PRIMARY KEY (id);


--
-- TOC entry 3650 (class 2606 OID 19212)
-- Name: item_categories pk_item_categories_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.item_categories
    ADD CONSTRAINT pk_item_categories_id PRIMARY KEY (id);


--
-- TOC entry 3655 (class 2606 OID 19228)
-- Name: items pk_items_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.items
    ADD CONSTRAINT pk_items_id PRIMARY KEY (id);


--
-- TOC entry 3648 (class 2606 OID 19661)
-- Name: npc_spawn_rules pk_npc_spawn_rules_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_spawn_rules
    ADD CONSTRAINT pk_npc_spawn_rules_id PRIMARY KEY (id);


--
-- TOC entry 3644 (class 2606 OID 19181)
-- Name: npc_templates pk_npc_templates_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_templates
    ADD CONSTRAINT pk_npc_templates_id PRIMARY KEY (id);


--
-- TOC entry 3659 (class 2606 OID 19316)
-- Name: races pk_races_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.races
    ADD CONSTRAINT pk_races_id PRIMARY KEY (id);


--
-- TOC entry 3642 (class 2606 OID 19652)
-- Name: world_node_features pk_world_node_features_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_node_features
    ADD CONSTRAINT pk_world_node_features_id PRIMARY KEY (id);


--
-- TOC entry 3639 (class 2606 OID 19155)
-- Name: world_nodes pk_world_nodes_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_nodes
    ADD CONSTRAINT pk_world_nodes_id PRIMARY KEY (id);


--
-- TOC entry 3635 (class 2606 OID 19263)
-- Name: worlds pk_worlds_id; Type: CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.worlds
    ADD CONSTRAINT pk_worlds_id PRIMARY KEY (id);


--
-- TOC entry 3677 (class 1259 OID 19693)
-- Name: idx_capabilities_capability_type; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_capabilities_capability_type ON game.capabilities USING btree (capability_type);


--
-- TOC entry 3678 (class 1259 OID 19688)
-- Name: idx_capabilities_parents; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_capabilities_parents ON game.capabilities USING btree (parent_id);


--
-- TOC entry 3679 (class 1259 OID 19689)
-- Name: idx_capabilities_tags_gin; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_capabilities_tags_gin ON game.capabilities USING gin (tags);


--
-- TOC entry 3680 (class 1259 OID 19691)
-- Name: idx_capabilities_type; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_capabilities_type ON game.capabilities USING btree (capability_type);


--
-- TOC entry 3681 (class 1259 OID 19687)
-- Name: idx_capabilities_world; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_capabilities_world ON game.capabilities USING btree (world_id);


--
-- TOC entry 3673 (class 1259 OID 19510)
-- Name: idx_character_class_features_class_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_character_class_features_class_id ON game.character_class_features USING btree (class_id);


--
-- TOC entry 3669 (class 1259 OID 19493)
-- Name: idx_character_classes_world_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_character_classes_world_id ON game.character_classes USING btree (world_id);


--
-- TOC entry 3670 (class 1259 OID 19492)
-- Name: idx_character_classes_world_id_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX idx_character_classes_world_id_code ON game.character_classes USING btree (world_id, code);


--
-- TOC entry 3665 (class 1259 OID 19452)
-- Name: idx_currency_world_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_currency_world_id ON game.currency USING btree (world_id);


--
-- TOC entry 3666 (class 1259 OID 19453)
-- Name: idx_currency_world_id_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX idx_currency_world_id_code ON game.currency USING btree (world_id, code);


--
-- TOC entry 3661 (class 1259 OID 19388)
-- Name: idx_enemies_world_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_enemies_world_id ON game.enemies USING btree (world_id);


--
-- TOC entry 3662 (class 1259 OID 19389)
-- Name: idx_enemies_world_id_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX idx_enemies_world_id_code ON game.enemies USING btree (world_id, code);


--
-- TOC entry 3651 (class 1259 OID 19237)
-- Name: idx_items_category_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_items_category_id ON game.items USING btree (category_id);


--
-- TOC entry 3652 (class 1259 OID 19369)
-- Name: idx_items_world; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_items_world ON game.items USING btree (world_id);


--
-- TOC entry 3645 (class 1259 OID 19235)
-- Name: idx_npc_spawn_rules_npc_template_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_npc_spawn_rules_npc_template_id ON game.npc_spawn_rules USING btree (npc_template_id);


--
-- TOC entry 3646 (class 1259 OID 19236)
-- Name: idx_npc_spawn_rules_world_node_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_npc_spawn_rules_world_node_id ON game.npc_spawn_rules USING btree (world_node_id);


--
-- TOC entry 3657 (class 1259 OID 19323)
-- Name: idx_races_world; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_races_world ON game.races USING btree (world_id);


--
-- TOC entry 3640 (class 1259 OID 19234)
-- Name: idx_world_node_features_world_node_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_world_node_features_world_node_id ON game.world_node_features USING btree (world_node_id);


--
-- TOC entry 3637 (class 1259 OID 19269)
-- Name: idx_world_nodes_world_id; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX idx_world_nodes_world_id ON game.world_nodes USING btree (world_id);


--
-- TOC entry 3653 (class 1259 OID 19360)
-- Name: ix_items_type; Type: INDEX; Schema: game; Owner: muddy
--

CREATE INDEX ix_items_type ON game.items USING btree (item_type);


--
-- TOC entry 3684 (class 1259 OID 19690)
-- Name: uidx_capabilities_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX uidx_capabilities_code ON game.capabilities USING btree (code);


--
-- TOC entry 3676 (class 1259 OID 19511)
-- Name: unq_character_class_features_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX unq_character_class_features_code ON game.character_class_features USING btree (code);


--
-- TOC entry 3656 (class 1259 OID 19368)
-- Name: uq_items_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX uq_items_code ON game.items USING btree (code);


--
-- TOC entry 3660 (class 1259 OID 19322)
-- Name: uq_races_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX uq_races_code ON game.races USING btree (world_id, code);


--
-- TOC entry 3636 (class 1259 OID 19370)
-- Name: uq_worlds_code; Type: INDEX; Schema: game; Owner: muddy
--

CREATE UNIQUE INDEX uq_worlds_code ON game.worlds USING btree (code);


--
-- TOC entry 3697 (class 2606 OID 19682)
-- Name: capabilities fk_capabilities_capabilities; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.capabilities
    ADD CONSTRAINT fk_capabilities_capabilities FOREIGN KEY (parent_id) REFERENCES game.capabilities(id) ON DELETE SET NULL;


--
-- TOC entry 3698 (class 2606 OID 19677)
-- Name: capabilities fk_capabilities_worlds; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.capabilities
    ADD CONSTRAINT fk_capabilities_worlds FOREIGN KEY (world_id) REFERENCES game.worlds(id) ON DELETE CASCADE;


--
-- TOC entry 3696 (class 2606 OID 19505)
-- Name: character_class_features fk_character_class_features_classes; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_class_features
    ADD CONSTRAINT fk_character_class_features_classes FOREIGN KEY (class_id) REFERENCES game.character_classes(id);


--
-- TOC entry 3695 (class 2606 OID 19487)
-- Name: character_classes fk_character_classes_worlds; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.character_classes
    ADD CONSTRAINT fk_character_classes_worlds FOREIGN KEY (world_id) REFERENCES game.worlds(id);


--
-- TOC entry 3694 (class 2606 OID 19447)
-- Name: currency fk_currency_worlds; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.currency
    ADD CONSTRAINT fk_currency_worlds FOREIGN KEY (world_id) REFERENCES game.worlds(id);


--
-- TOC entry 3693 (class 2606 OID 19383)
-- Name: enemies fk_enemies_worlds; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.enemies
    ADD CONSTRAINT fk_enemies_worlds FOREIGN KEY (world_id) REFERENCES game.worlds(id);


--
-- TOC entry 3689 (class 2606 OID 19213)
-- Name: item_categories fk_item_categories_parent_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.item_categories
    ADD CONSTRAINT fk_item_categories_parent_id FOREIGN KEY (parent_id) REFERENCES game.item_categories(id);


--
-- TOC entry 3690 (class 2606 OID 19229)
-- Name: items fk_items_category_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.items
    ADD CONSTRAINT fk_items_category_id FOREIGN KEY (category_id) REFERENCES game.item_categories(id);


--
-- TOC entry 3691 (class 2606 OID 19363)
-- Name: items fk_items_world_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.items
    ADD CONSTRAINT fk_items_world_id FOREIGN KEY (world_id) REFERENCES game.worlds(id);


--
-- TOC entry 3687 (class 2606 OID 19192)
-- Name: npc_spawn_rules fk_npc_spawn_rules_npc_template_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_spawn_rules
    ADD CONSTRAINT fk_npc_spawn_rules_npc_template_id FOREIGN KEY (npc_template_id) REFERENCES game.npc_templates(id);


--
-- TOC entry 3688 (class 2606 OID 19197)
-- Name: npc_spawn_rules fk_npc_spawn_rules_world_node_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.npc_spawn_rules
    ADD CONSTRAINT fk_npc_spawn_rules_world_node_id FOREIGN KEY (world_node_id) REFERENCES game.world_nodes(id);


--
-- TOC entry 3692 (class 2606 OID 19317)
-- Name: races fk_races_world_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.races
    ADD CONSTRAINT fk_races_world_id FOREIGN KEY (world_id) REFERENCES game.worlds(id);


--
-- TOC entry 3686 (class 2606 OID 19166)
-- Name: world_node_features fk_world_node_features_world_node_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_node_features
    ADD CONSTRAINT fk_world_node_features_world_node_id FOREIGN KEY (world_node_id) REFERENCES game.world_nodes(id);


--
-- TOC entry 3685 (class 2606 OID 19264)
-- Name: world_nodes fk_world_nodes_worlds_id; Type: FK CONSTRAINT; Schema: game; Owner: muddy
--

ALTER TABLE ONLY game.world_nodes
    ADD CONSTRAINT fk_world_nodes_worlds_id FOREIGN KEY (world_id) REFERENCES game.worlds(id);


-- Completed on 2025-05-28 18:47:48 EDT

--
-- PostgreSQL database dump complete
--

