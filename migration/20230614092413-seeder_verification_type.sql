-- +migrate Up
INSERT INTO public.verification_type (id, code, name, "desc", is_active_flag, created_at, created_by, updated_at, updated_by, deleted_at) VALUES (DEFAULT, 'VT-0001', 'OTP Verification', 'Verification OTP', true, '2023-06-14 09:34:20.000000', 0, '2023-06-14 09:34:27.000000', 0, null);

INSERT INTO public.verification_type (id, code, name, "desc", is_active_flag, created_at, created_by, updated_at, updated_by, deleted_at) VALUES (DEFAULT, 'VT-0002', 'OTP Verification Reset Password', 'Verification OTP Reset Password', true, '2023-06-14 09:34:20.000000', 0, '2023-06-14 09:34:27.000000', 0, null);

-- +migrate Down
