--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

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
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.failed_jobs (
    id bigint NOT NULL,
    uuid character varying(255) NOT NULL,
    connection text NOT NULL,
    queue text NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.failed_jobs OWNER TO postgres;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.failed_jobs_id_seq OWNER TO postgres;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- Name: foods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foods (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    stock bigint NOT NULL,
    price bigint NOT NULL,
    deleted_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.foods OWNER TO postgres;

--
-- Name: migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.migrations (
    id integer NOT NULL,
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE public.migrations OWNER TO postgres;

--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.migrations_id_seq OWNER TO postgres;

--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    food_id uuid NOT NULL,
    quantity bigint NOT NULL,
    price bigint NOT NULL,
    deleted_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: password_resets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.password_resets (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE public.password_resets OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    email_verified_at timestamp(0) without time zone,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    deleted_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Data for Name: failed_jobs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.failed_jobs (id, uuid, connection, queue, payload, exception, failed_at) FROM stdin;
\.


--
-- Data for Name: foods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.foods (id, name, description, stock, price, deleted_at, created_at, updated_at) FROM stdin;
9d38abf5-5815-4439-9d2a-95f3623990a2	Amet rerum dicta nam ipsam sint.	Voluptatum dolor consequuntur blanditiis dolores dolorem pariatur. Aliquam vitae tenetur rerum vero aliquid iure. Et ullam dignissimos est suscipit accusantium. Est nam et pariatur placeat. Et cumque dolor voluptatum autem esse dolore repellat.	6	8007	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
cfc51f28-7792-4266-a7d3-f28ed2b937bc	Voluptas maxime laborum reiciendis voluptas aliquam.	Dolor modi neque est consequatur odio quos similique tempore. Eius et fugiat expedita quis. Pariatur est voluptates sed repellat culpa totam officiis et.	4	3709	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
afcfd620-d4cd-436e-9dc7-04807a6fe035	Est nulla est vel ipsa assumenda vel.	Voluptatem delectus perferendis velit repellendus deserunt officiis quis. Qui quia delectus veritatis id amet necessitatibus quo et. Perferendis et ducimus ipsum repudiandae asperiores quia omnis. Eaque voluptatibus labore repellat quisquam doloribus beatae.	4	1579	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
fbdffb11-995a-4925-a29c-0414deea2602	Maxime illum a et aliquid veniam voluptas.	Aliquid dolorem et quos iure eligendi non repellendus. Adipisci totam sunt fugit ea. Quos expedita in quia officiis. Maxime ipsa itaque labore esse. Voluptas quam facilis quis culpa magni laboriosam.	6	8879	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
dcd837a5-7edf-44de-9fdc-e65a9a7bffa9	Modi ea quis earum ipsa unde incidunt.	Autem deleniti eveniet incidunt consequatur. Optio deleniti incidunt quos mollitia.	2	2440	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
e68ce0ca-7e09-465e-b332-29218e4c818f	Ipsam ea enim et vel.	Corrupti ipsum quisquam ad soluta libero tempore rerum. Voluptas est nam quisquam id qui quo. Sed consequuntur odit molestiae recusandae.	1	2581	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
1babc394-c3d0-40ae-a568-9b0634bf248d	Et placeat quibusdam voluptatum.	Sed ea quibusdam hic voluptatem et atque. Consequatur repellendus autem et. Et qui nulla in iure iste. Unde quia illo non.	1	3458	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
e5d9022f-c091-4c69-9923-383324946985	Pariatur assumenda corporis eum nobis unde fuga unde.	Rerum et aliquid occaecati. Vel voluptatem qui accusantium optio enim iste. Totam eaque a laudantium iste. Quasi nisi ab illum voluptatem debitis similique.	3	3755	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
b127ab6c-7c0e-4dad-88f0-fc031e1157f7	Maiores dolore earum dolor ut modi et molestiae.	Voluptates voluptatem ea inventore tempore facilis assumenda. Excepturi expedita quia necessitatibus facilis a voluptatem. Aut modi id eos possimus ut. Ducimus hic animi non dignissimos fuga cupiditate tempore sit.	2	4418	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
fe6c7137-2b00-48d4-aa98-181021b816dc	Eius ipsum possimus nihil laudantium ut veniam.	Quo accusamus molestiae illum. Odio et illum autem est ipsam adipisci. Eos possimus sequi ut eum alias autem aut. Quidem ipsa pariatur nemo ratione voluptatem iste.	7	7895	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
1a0466c3-fe9c-437d-9ece-fa6a19f39602	Itaque ut ipsum optio asperiores qui assumenda doloremque fugit.	Eum ipsa enim ratione. Illo eos nulla et accusamus. Sit laboriosam quis voluptas veritatis earum necessitatibus sequi.	0	3077	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
dc80629e-2afa-4fa3-aade-c1eeaaccfa5b	Ut nesciunt ratione repellat asperiores dicta saepe.	Rerum distinctio quasi quas nisi. Facere quis aut molestiae dolor occaecati et. Delectus nam non dolor. Eum aut maiores placeat eos eos sint.	3	6562	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
9cc4f526-e7d3-4f24-aa4e-b839ae07b1ee	Ipsam magni consequatur adipisci et suscipit nostrum blanditiis.	Doloribus voluptatem cum facilis vel optio reiciendis ipsa qui. Aut nulla est a animi unde. Voluptatem nihil libero voluptas. Optio qui optio voluptatem sint et reiciendis.	6	5832	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
6539d458-1434-489d-a12e-1202037f1467	Itaque velit eaque sit saepe nulla.	Exercitationem at fuga in. Ipsa ea minus omnis qui et. Aut neque occaecati expedita id. Perferendis tempora libero atque eaque ipsum quas fugiat.	6	1484	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
e0aa6a88-34b6-40dd-a068-aed2c51c6f81	Consectetur est ut cum est dignissimos voluptates reiciendis.	Doloremque dolorum et et ut occaecati neque. Eaque laborum est id nihil voluptatem saepe aut. Ex atque laborum sed ullam ut recusandae. Quas tempore quidem asperiores consequatur accusamus repellat accusantium.	0	8039	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
0c6d6a39-5866-421a-9e31-26415c6e41a7	Dolor eius itaque blanditiis autem velit.	Corrupti dolorem quia ad et. Similique quae quibusdam deleniti quos eum vitae. Praesentium sunt ipsum ullam aspernatur esse dolore inventore. Voluptas qui consectetur numquam corrupti.	6	5315	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
e5451cb7-3ee6-4f4c-99a5-838a8ebc96e5	Laboriosam qui aspernatur aspernatur numquam.	Esse porro corporis ipsum quisquam iure aut aut. Consequatur laborum aut molestiae consequatur qui. Aperiam illum laudantium quam quae esse quidem.	6	4934	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
0db4df23-f790-405f-8ed6-11b4367e990a	Ratione ad error et recusandae.	Voluptatibus eius est reiciendis consequatur qui tempora. Magni voluptas odit blanditiis harum. Debitis aut non possimus qui omnis. Voluptatem labore et est dolor cupiditate excepturi debitis.	0	7987	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
4c4e3d6f-2d3e-4c78-bcda-fb20499a4561	Reprehenderit dolorum aliquam odio commodi.	A distinctio omnis et. Excepturi non incidunt ratione atque quae. Qui dolorem ut numquam vitae consequatur velit.	9	4188	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
346935c2-e1ff-48f1-8801-e15679fbb4fb	Quis accusamus aliquid asperiores quibusdam aliquam hic culpa.	Esse corrupti iure recusandae tenetur quia est neque. Minima pariatur nobis molestiae aut. Repellat sed dolores dolor sed assumenda corporis eius. Et et eos ut magnam numquam. Nam est aut dolor nostrum nihil et minima.	8	8700	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
58d78434-6635-49d8-98e6-792cbf670a7e	Quo qui eius ut est rem officia aut.	Et quis voluptatibus facilis voluptatem labore ut. Est voluptatem et autem fugiat dolorem et. Nulla sit delectus quos autem.	6	2196	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
51754e86-7a45-4566-81d2-0d0016894eb6	Et molestiae tempora maiores.	Sequi error minima quibusdam inventore. Sed quas molestiae omnis odit impedit minus recusandae. Aut dolorum sapiente facilis dolores est ab voluptatibus sit. Voluptatibus fugit laudantium ut earum commodi autem.	0	4328	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
7668b7fe-e530-4aed-a195-2fbb8330c87a	Odio magni perspiciatis voluptas adipisci dignissimos nemo.	Animi et recusandae dolor sed sit. Quos quis sint maxime similique repudiandae neque. Fuga eveniet nisi ab. Suscipit perspiciatis sed harum distinctio inventore.	8	4995	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
cc78e777-4991-49db-9e52-530df2b64f0b	Ipsam non consequatur assumenda.	Dolorem dolorem fugiat et. Molestias voluptatem et odit fugiat magnam repudiandae molestiae. Eum et quos autem id ut autem provident consequatur. Aut vel eum debitis ab quisquam ea velit.	5	6861	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
338e949e-1043-4e0f-8238-681d1c1a4322	Delectus quisquam ut voluptatum blanditiis ut eligendi facere non.	Voluptatem qui reiciendis voluptate nam soluta voluptates deserunt. Illo inventore quasi veniam ut et et. Sed voluptatibus quidem ducimus maiores voluptas facilis rerum.	5	2962	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
0d2da4c5-5c95-4ecf-a183-5921f5828481	Sunt repellat unde quia vel.	Facilis optio adipisci veritatis perferendis et delectus. Ducimus quae vitae similique accusamus ratione. Officiis quis aliquid explicabo vero libero. Odio magnam velit corrupti aut et.	1	6788	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
a9528a3f-de53-45fa-a606-f1808870fa44	Voluptatibus voluptatem ullam doloremque aperiam vitae aut.	Consectetur perspiciatis voluptatem sapiente doloremque quam totam. Velit ut dolor sapiente. Totam ab aut nesciunt nostrum. Laborum dolores dignissimos repudiandae explicabo eligendi eum impedit.	5	7720	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
98756e37-7939-40f7-a767-27cb3d121302	Quia ut optio quis impedit quae sint in.	Veritatis assumenda et aut omnis ut temporibus quos. Vitae modi asperiores et laudantium debitis nemo. Sunt dolor nulla suscipit est dicta est. Iusto ea dolorem tempora omnis aspernatur modi dolorum.	2	3689	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
f5d0736f-7373-42e6-9b0a-536d80006195	Id ea vel exercitationem animi.	Quia omnis qui tempora dolorem rerum. Aut ipsa iusto tenetur dolores quia. Quasi autem praesentium delectus explicabo vitae dolorum.	3	2414	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
14cc253d-556b-4786-91b5-90b96aed7742	Optio quia fuga ad.	Fuga qui voluptatem et error. Facilis soluta facilis reiciendis blanditiis accusamus aliquid ipsam. Delectus illo facere est. Autem quasi veniam facilis. Dicta dolores optio ut delectus.	9	1701	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
fefd50cf-12eb-4674-8d5a-404a643a23ee	Quae nobis nobis tempora distinctio ipsa sint.	Accusamus expedita nesciunt ut quia quaerat dolores et. Mollitia nostrum officiis ducimus accusamus qui nesciunt fugiat. Magni quisquam cum quidem. Voluptas sunt et repudiandae placeat non est aut.	3	2243	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
66834d35-1fb0-47ff-903c-677fe343ee1d	Veniam et natus dolor sequi unde deleniti.	Repellat esse dolore aut ea est reprehenderit. Pariatur omnis sapiente eius ea voluptatum esse. Ullam velit nemo quod doloremque minus temporibus libero.	1	2020	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
aa45ff3a-4344-48fa-bd30-41c194d2c5b3	Debitis numquam molestiae eveniet adipisci maiores.	Tempore perferendis ipsa velit temporibus. Possimus velit aut sed. Rerum provident quas rerum. Rerum consequatur nulla dicta qui et iste non.	2	6123	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
50a69370-76dd-4552-a327-b0420f3bf33e	Quo nesciunt voluptas rem natus itaque quis eaque ut.	Optio cum quibusdam aspernatur neque neque ex nihil. Ut commodi debitis dolorem. Et laboriosam aut vitae quae vero. Recusandae autem nemo qui dolorum.	2	7440	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
abefcfcb-2c01-4ca1-af80-affe7e4407e8	Laborum in eum dignissimos ab fugit consectetur sint.	Similique beatae dolor beatae voluptatem voluptatum accusantium quo. Nulla et et magnam dolor est et. Laudantium accusantium in enim itaque.	9	8548	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
536755b7-2442-4a5b-aeb6-d314bdd7ba3c	Quaerat earum aliquid voluptas quo.	Quae totam voluptas consectetur quis natus eveniet consectetur. Et rerum consequatur autem reiciendis exercitationem. Sunt autem provident ratione rerum voluptates.	2	1791	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
5a1a40b3-3b02-42f3-b429-ce2ed5a5563c	Nemo soluta a dolorem voluptas.	Voluptate corporis blanditiis voluptates reprehenderit consequatur omnis. Sed sapiente architecto mollitia ex aut perferendis. Velit laudantium earum error omnis officiis adipisci. Perferendis laborum iure autem nulla.	8	4986	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
bdfd9d2d-f84c-4a96-a692-dbb99ad1e9e0	Sunt porro minima iusto et ad.	Harum harum illo temporibus ratione ut beatae. Quam quibusdam sit illum dolorem. Non qui nostrum numquam repudiandae laudantium aspernatur maiores ut. Ut sed aliquam veritatis sunt est.	5	4182	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
640d6e7d-6813-4354-bbd0-adfef4c346af	Provident ipsam facilis eligendi.	Illum et nisi quibusdam ea quis aut maxime. Pariatur exercitationem non dignissimos quo. Delectus iure nobis officia eius.	3	1164	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
6c3abca0-a016-42ab-b597-77b579d8fb51	Consequatur aut assumenda nihil officia nobis qui tenetur ducimus.	Repudiandae qui deserunt aut est. Ipsa repellendus qui quae voluptas ex error.	3	7399	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
395cab6e-f9ed-44ed-8517-1e44347ee001	Laboriosam dolor molestiae consequatur est dolorem et eveniet.	Ut illum provident rerum aut ut. Et incidunt aliquam atque iure dignissimos laudantium voluptatem. Est repudiandae pariatur ut ipsa et optio consectetur.	7	7868	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
a110cf4f-4d44-4458-8703-0e8394ed2463	Voluptate dolores ipsa necessitatibus nesciunt dignissimos placeat et.	Rem quasi dolores aperiam similique voluptatem ducimus ipsa. Officia architecto ut consequatur aliquam qui. Quia et architecto perspiciatis dolorem maxime assumenda.	4	7512	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
cbcbcb33-3d32-41af-9e7c-4bdad9f44fc0	Eos omnis magnam sit natus.	Omnis autem voluptatum nobis quidem et dolores. Aut hic eos aliquid corporis. Molestiae impedit omnis repellat culpa voluptatum.	1	5826	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
54178370-a5fb-4d75-901a-d74bd32cfedb	Esse qui rem aperiam velit.	Quae sint nobis quibusdam asperiores voluptates placeat. Illo adipisci reiciendis velit itaque ut ullam eum accusantium. Beatae ducimus cupiditate qui quibusdam quae. Possimus dolores et eos alias.	3	6802	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
a7415fc3-bb78-4643-82b1-fc966e91416c	Necessitatibus voluptates pariatur ea aut velit.	Voluptatum harum impedit provident minima sapiente. Qui porro ut et reprehenderit. Cupiditate iure numquam ea doloribus natus consequuntur. Magni vitae corporis sed accusantium qui.	2	5027	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
bbf76a88-03a9-4001-b149-c5894326b65b	Suscipit dicta repellat consectetur voluptate nihil aspernatur.	Excepturi aut rerum et. Qui quas accusamus et ratione quis. Rerum autem sed officia. Vero odio aut quo consectetur temporibus sunt architecto officiis. Placeat distinctio laborum libero sed facere vitae voluptates.	4	4285	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
6e1e98ec-027e-49b3-afb3-b8ee9da52ccb	Dolorem error quis quia molestiae voluptas error sunt.	Ipsa autem et numquam nihil alias et. Asperiores suscipit et veniam mollitia molestias saepe magni. Sequi voluptatum illum enim aut. Consequatur at tempora accusamus.	8	2169	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
9d3f7a0b-7eec-4d5f-986e-f25bff0a834c	Iste ea iusto ea sint accusantium distinctio reiciendis.	Laborum qui ipsam nihil cumque ullam vero. Consequatur quidem maiores qui minus. Unde qui culpa cumque dolore natus. Et non non quos similique quam ab.	9	7446	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
3859ca5e-359d-433d-9a85-22c5152a9392	Quod tenetur quasi excepturi sed sit iure.	Excepturi aut maxime rerum dolor dolore. Unde fuga incidunt sint aperiam distinctio iste. Quis et voluptatem ut quasi. Et quo odit quo sed omnis consequuntur.	9	2626	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
5999c46c-d197-48b0-b1f3-330ac4eb9445	Rerum vel illum aliquid sit aut et maxime.	Nulla tempore nulla qui illum commodi sit non. Et velit sequi tempora. Deserunt pariatur et perspiciatis suscipit voluptatem optio modi in.	5	2146	\N	2020-12-14 08:02:40	2020-12-14 08:02:40
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.migrations (id, migration, batch) FROM stdin;
1	2014_10_12_000000_create_users_table	1
2	2014_10_12_100000_create_password_resets_table	1
3	2019_08_19_000000_create_failed_jobs_table	1
4	2020_12_13_075516_create_foods_table	1
5	2020_12_13_075936_create_orders_table	1
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, user_id, food_id, quantity, price, deleted_at, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: password_resets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.password_resets (email, token, created_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, type, email, email_verified_at, password, remember_token, deleted_at, created_at, updated_at) FROM stdin;
91381e6f-86a1-42b8-b2fb-cc83f0152760	Prof. Mayra Kassulke III	user	zernser@example.org	2020-12-14 08:02:39	$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi	nkY5QJlAjM	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
c150a721-4ce4-4ba7-9727-c731774098b9	Kaley Collins	admin	bode.isom@example.org	2020-12-14 08:02:39	$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi	89JhhmdIV3	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
73d6a795-ed80-460a-8b2e-aa278a32030a	Olga Kertzmann IV	user	wilford.mcdermott@example.net	2020-12-14 08:02:39	$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi	dE2bGWZjwN	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
9ac00ec7-e249-451c-b84c-cc2119a34d6c	Mr. Hal Collier	user	smith.delia@example.org	2020-12-14 08:02:39	$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi	UUd0eHpsW4	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
8fddfc5f-4f26-413b-8bf8-c05e50860e30	Abigail Mann	user	olson.janae@example.net	2020-12-14 08:02:39	$2a$10$nRkEUcZQM7ii8QBebmHUX.vpqtBVC.CR0ZN9/RXfdXj5VWn733qqe	EpvFGdEDtm	\N	2020-12-14 08:02:39	2020-12-14 08:02:39
\.


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.failed_jobs_id_seq', 1, false);


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.migrations_id_seq', 5, true);


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- Name: failed_jobs failed_jobs_uuid_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_uuid_unique UNIQUE (uuid);


--
-- Name: foods foods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foods
    ADD CONSTRAINT foods_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: password_resets_email_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX password_resets_email_index ON public.password_resets USING btree (email);


--
-- Name: orders orders_food_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_food_id_foreign FOREIGN KEY (food_id) REFERENCES public.foods(id);


--
-- Name: orders orders_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

