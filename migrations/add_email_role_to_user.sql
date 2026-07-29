-- Migration: tambah kolom Email dan Role ke tabel user
-- Jalankan script ini di database MySQL Anda

-- Tambah kolom Email (unik, boleh NULL dulu untuk kompatibilitas data lama)
ALTER TABLE `user`
  ADD COLUMN IF NOT EXISTS `Email` VARCHAR(255) NULL AFTER `Username`,
  ADD COLUMN IF NOT EXISTS `Role`  VARCHAR(50)  NOT NULL DEFAULT 'user' AFTER `Password`;

-- Tambah UNIQUE constraint setelah kolom ada
-- (skip jika sudah ada)
ALTER TABLE `user`
  ADD UNIQUE INDEX IF NOT EXISTS `uni_user_username` (`Username`),
  ADD UNIQUE INDEX IF NOT EXISTS `uni_user_email` (`Email`);

-- Verifikasi hasil migrasi
DESCRIBE `user`;
