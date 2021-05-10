-- INSERT APPLICATION SERVICE
-- INSERT AUTHENTICATE SERVICE
INSERT INTO public.apps(uuid, name, latest_version, running_status, type)
VALUES  ('2fe929fb-8f2a-4f16-9b54-51a4dbc61daf','Root Application Service', '1.0', 'active', 'service'),
        ('2fe929fb-8f2a-4f16-9b54-51a4dbc61frs','Authenticate Service', '1.0', 'active', 'service');



INSERT INTO public.configs(log_mode, migration_db, server_host, port, access_secret, refresh_secret, access_duration, refresh_duration, app_uuid)
VALUES  ('yes', 'yes', '127.0.0.1', '8000', 'gRPC-app-acc', 'gRPC-app-ref', '3600', '86400', '2fe929fb-8f2a-4f16-9b54-51a4dbc61daf'),
        ('yes', 'yes', '127.0.0.1', '8000', 'gRPC-app-acc', 'gRPC-app-ref', '3600', '86400', '2fe929fb-8f2a-4f16-9b54-51a4dbc61daf');


INSERT INTO public.firewalls(host, port, app_uuid)
VALUES  ('*', '*', '*', '2fe929fb-8f2a-4f16-9b54-51a4dbc61daf'),
        ('*', '*', '*', '2fe929fb-8f2a-4f16-9b54-51a4dbc61frs');

INSERT INTO public.databases(type, name, connection_string, app_uuid)
VALUES ('postgres', 'gRPC-FSA-Application', 'postgres://%v:%v@%v:%v/%v?sslmode=disable')