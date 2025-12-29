-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 29 Des 2025 pada 06.07
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
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `news`
--

INSERT INTO `news` (`id`, `title`, `content`, `created_at`) VALUES
(1, 'Berita Fiber', 'Ini contoh REST API menggunakan Fiber', '2025-12-23 05:39:41'),
(2, 'Berita kelahiran', 'Ini dalah berita kelahiran sang juruslamat', '2025-12-23 06:03:50'),
(4, 'Berita kelahiran', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-23 08:14:51'),
(5, 'Berita kelahiran', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-23 09:19:16'),
(6, 'Berita kelahirantttnn', 'Ini adalah berita kelahiran sang juruslamat', '2025-12-27 01:05:48'),
(7, 'Berita ksdfgersgeselamatgykjmgfjftannn', 'kehidupan baru', '2025-12-27 01:31:56'),
(8, 'Judul updatehgfyffhh', 'Isi update1', '2025-12-27 01:53:08'),
(9, 'Berita keselamatan', 'kehidupan baru', '2025-12-27 01:58:09'),
(10, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00'),
(11, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00'),
(12, 'Berita kehadiran pak prabowo', 'bantuan untuk saudara aceh', '0000-00-00 00:00:00'),
(16, 'p', 'p', '2025-12-28 20:37:03');

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
