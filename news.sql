-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 29 Des 2025 pada 10.54
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `newsapp`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `news`
--

CREATE TABLE `news` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `author` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `news`
--

INSERT INTO `news` (`id`, `title`, `content`, `created_at`, `updated_at`, `author`) VALUES
(1, 'Berita Fiber', 'Ini contoh REST API menggunakan Fiber', '2025-12-23 05:39:41', NULL, ''),
(2, 'Berita kelahiran', 'Ini dalah berita kelahiran sang juruslamat', '2025-12-23 06:03:50', NULL, ''),
(4, 'Berita kelahiran', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-23 08:14:51', NULL, ''),
(5, 'Berita kelahiran', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-23 09:19:16', NULL, ''),
(6, 'Berita kelahirantttnn', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-27 01:05:48', NULL, ''),
(7, 'Berita ksdfgersgeselamatgykjmgfjftannn', 'kehidupan baru', '2025-12-27 01:31:56', NULL, ''),
(8, 'berita hari ini', 'selamat datang pada hari ke 4 belajar', '2025-12-27 01:53:08', '2025-12-29 07:56:01', ''),
(9, 'Berita keselamatan', 'kehidupan baru', '2025-12-27 01:58:09', NULL, ''),
(10, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00', NULL, ''),
(11, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00', NULL, ''),
(12, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00', NULL, ''),
(16, 'asfasf', 'kesdfhidusdfsdfpan baru', '2025-12-28 20:37:03', NULL, ''),
(17, 'prege', 'pge', '2025-12-28 22:36:07', NULL, ''),
(18, 'prege', 'pgtjhrre', '2025-12-28 22:46:11', NULL, ''),
(19, 'prege', 'pgtjhrre', '2025-12-28 22:56:30', NULL, ''),
(20, 'prege', 'pgtjhrre', '2025-12-28 23:22:02', NULL, ''),
(21, 'prege', 'pgtjhrre', '2025-12-28 23:27:01', NULL, ''),
(22, 'prege', 'pgtjhrre', '2025-12-28 23:27:56', NULL, ''),
(23, 'prege', 'pgtjhrre', '2025-12-28 23:42:20', NULL, ''),
(24, 'prege', 'pasfbasjkbfkasnafasasasgtjhrre', '2025-12-29 00:36:28', '2025-12-29 07:36:28', ''),
(25, 'holla', 'hhhaiiii', '2025-12-29 00:40:18', '2025-12-29 07:40:54', ''),
(26, 'mikroskil', 'kampus', '2025-12-29 00:42:13', '2025-12-29 07:42:13', ''),
(27, 'holla', 'hhhaiiii', '2025-12-29 00:43:01', '2025-12-29 07:44:14', ''),
(28, 'berita hari ini', 'selamat datang pada hari ke 4 belajar', '2025-12-29 01:45:48', '2025-12-29 08:46:40', '');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `news`
--
ALTER TABLE `news`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `news`
--
ALTER TABLE `news`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
