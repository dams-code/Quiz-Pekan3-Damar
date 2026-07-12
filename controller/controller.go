package controller

import (
	"fmt"
	"net/http"
	"quiz-pekan3-damar/middleware"
	"quiz-pekan3-damar/repository"
	"quiz-pekan3-damar/structbuku"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {
	var setUser structbuku.Users
	if err := ctx.ShouldBindJSON(&setUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_login":  "error",
			"status_detail": err.Error(),
		})
		return
	}

	userFromDB, err := repository.GetUsername(DBSqlConn, setUser.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status_login":  "error",
			"status_detail": "Username atau password salah",
		})
		return
	}

	if userFromDB.Password != setUser.Password {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status_login":  "error",
			"status_detail": "Username atau password salah",
		})
		return
	}

	token, err := middleware.GenerateJWT(userFromDB.ID, userFromDB.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_login":  "error",
			"status_detail": "Gagal generate token: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_login": "sukses",
		"message":      "Hi, " + userFromDB.Username,
		"token":        token,
	})
}

func TambahUser(ctx *gin.Context) {
	var setUser structbuku.Users
	if err := ctx.ShouldBindJSON(&setUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_tambah_user": "error",
			"error":              err.Error(),
		})
		return
	}

	HasilTambahUser, err := repository.TambahUser(DBSqlConn, setUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_tambah_user": "error",
			"status_detail":      err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status_tambah_user": "sukses",
		"data": gin.H{
			"id":         HasilTambahUser.ID,
			"username":   HasilTambahUser.Username,
			"password":   "****",
			"created_at": HasilTambahUser.CreatedAt,
		},
	})
}

func GetBuku(ctx *gin.Context) {
	hasilGetBuku, err := repository.GetBuku(DBSqlConn)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_get_All_Buku": "error",
			"status_detail":       err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_get_All_Buku": "sukses",
		"data":                hasilGetBuku,
	})
}

func GetBukuID(ctx *gin.Context) {
	setIdBuku, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_get_bukuId": "error",
			"status_detail":     "Id buku harus angka",
		})
		return
	}

	hasilGetBukuID, err := repository.GetBukuID(DBSqlConn, setIdBuku)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_get_bukuId": "error",
			"status_detail":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_get_bukuId": "sukses",
		"data":              hasilGetBukuID,
	})
}

func GetBukuBerdasarkanKategori(ctx *gin.Context) {
	setIdKategoriBuku, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_get_kategoriBukuId": "error",
			"status_detail":             "Id Kategori harus angka",
		})
		return
	}

	HasilGetBukuBerdasarkanKategori, err := repository.GetBukuBerdasarkanKategori(DBSqlConn, setIdKategoriBuku)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_get_kategoriBukuId": "error",
			"status_detail":             err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_get_kategoriBukuId": "sukses",
		"data":                      HasilGetBukuBerdasarkanKategori,
	})
}

func TambahBuku(ctx *gin.Context) {
	var setBuku structbuku.Buku

	if err := ctx.ShouldBindJSON(&setBuku); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_tambah_buku": "error",
			"status_detail":      err.Error(),
		})
		return
	}

	if setBuku.ReleaseYear == nil || *setBuku.ReleaseYear < 1980 || *setBuku.ReleaseYear > 2024 {
		ctx.JSON(400, gin.H{"status_tambah_buku": "Tahun rilis buku harus antara 1980 dan 2024."})
		return
	}

	usernameJWT, exists := ctx.Get("username_sekarang")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Gagal mengidentifikasi user dari token"})
		return
	}

	usernameStr := usernameJWT.(string)

	setBuku.CreatedBy = &usernameStr

	HasilTambahBuku, err := repository.TambahBuku(DBSqlConn, setBuku)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_tambah_buku": "error",
			"status_detail":      err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status_tambah_buku": "sukses",
		"data":               HasilTambahBuku,
	})
}

func HapusBuku(ctx *gin.Context) {
	setIdBuku, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_hapus_buku": "error",
			"status_detail":     "Id Buku harus angka",
		})
		return
	}

	err = repository.HapusBuku(DBSqlConn, setIdBuku)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_hapus_buku": "error",
			"status_detail":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_hapus_buku": "sukses",
		"message":           fmt.Sprintf("Buku ID %d berhasil dihapus", setIdBuku),
	})
}

func UpdateBuku(ctx *gin.Context) {

	var setBuku structbuku.Buku

	if err := ctx.ShouldBindJSON(&setBuku); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_update_buku": "error",
			"status_detail":      err.Error(),
		})
		return
	}

	setIdBuku, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_update_buku": "error",
			"status_detail":      "Id Buku harus angka",
		})
		return
	}

	HasilUpdateBuku, err := repository.UpdateBuku(DBSqlConn, setIdBuku, setBuku)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_update_buku": "error",
			"status_detail":      err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_update_buku": "sukses",
		"data":               HasilUpdateBuku,
	})
}

func GetKategori(ctx *gin.Context) {
	hasilGetKategori, err := repository.GetKategori(DBSqlConn)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_get_All_Kategori": "error",
			"status_detail":           err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_get_All_Kategori": "sukses",
		"data":                    hasilGetKategori,
	})
}

func GetKategoriID(ctx *gin.Context) {
	setIdKategori, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_get_kategoriId": "error",
			"status_detail":         "Id kategori harus angka",
		})
		return
	}

	hasilGetkategoriID, err := repository.GetKategoriID(DBSqlConn, setIdKategori)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_get_kategoriId": "error",
			"status_detail":         err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_get_kategoriId": "sukses",
		"data":                  hasilGetkategoriID,
	})
}

func TambahKategori(ctx *gin.Context) {
	var setKategori structbuku.Kategori

	if err := ctx.ShouldBindJSON(&setKategori); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_tambah_kategori": "error",
			"status_detail":          err.Error(),
		})
		return
	}

	usernameJWT, exists := ctx.Get("username_sekarang")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Gagal mengidentifikasi user dari token"})
		return
	}

	usernameStr := usernameJWT.(string)

	setKategori.CreatedBy = usernameStr

	HasilTambahKategori, err := repository.TambahKategori(DBSqlConn, setKategori)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_tambah_kategori": "error",
			"status_detail":          err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status_tambah_kategori": "sukses",
		"data":                   HasilTambahKategori,
	})
}

func HapusKategori(ctx *gin.Context) {
	setIdKategori, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_hapus_kategori": "error",
			"status_detail":         "Id Kategori harus angka",
		})
		return
	}

	err = repository.HapusKategori(DBSqlConn, setIdKategori)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_hapus_kategori": "error",
			"status_detail":         err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_hapus_kategori": "sukses",
		"message":               fmt.Sprintf("Kategori dengan ID %d berhasil dihapus", setIdKategori),
	})
}

func UpdateKategori(ctx *gin.Context) {
	setIdKategori, err := strconv.Atoi(ctx.Param("id"))

	username, exists := ctx.Get("username_sekarang")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status_update_kategori": "error",
			"status_detail":          "User tidak terdeteksi",
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_update_kategori": "error",
			"status_detail":          "Id Kategori harus angka",
		})
		return
	}

	var setKategori structbuku.Kategori

	setUsername := username.(string)

	setKategori.ModifiedBy = &setUsername

	if err := ctx.ShouldBindJSON(&setKategori); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_update_kategori": "error",
			"status_detail":          err.Error(),
		})
		return
	}

	HasilUpdateKategori, err := repository.UpdateKategori(DBSqlConn, setIdKategori, setKategori)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_update_kategori": "error",
			"status_detail":          err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_update_kategori": "sukses",
		"data":                   HasilUpdateKategori,
	})
}
