package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/models"
)

func ObtenerRegistrosFinanzas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_finanzas, id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo, fecha_creacion, fecha_modificacion FROM finanzas ORDER BY id_finanzas`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar finanzas")
		return
	}
	defer rows.Close()
	items := []models.Finanzas{}
	for rows.Next() {
		item, err := scanFinanzas(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer finanzas")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerFinanzasPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_finanzas, id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo, fecha_creacion, fecha_modificacion FROM finanzas WHERE id_finanzas = $1`, id)
	item, err := scanFinanzas(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "finanzas no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar finanzas")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearFinanzas(w http.ResponseWriter, r *http.Request) {
	var item models.Finanzas
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesFinanzas(item.IdUsuario, item.IdMovimientoDinero, item.IdCategoria); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`INSERT INTO finanzas (id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id_finanzas, id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdMovimientoDinero, item.IdCategoria, item.MontoPresupuesto, item.Gasto, item.Disponible, item.Fecha, item.Activo)
	item, err := scanFinanzas(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear finanzas")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarFinanzas(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Finanzas
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesFinanzas(item.IdUsuario, item.IdMovimientoDinero, item.IdCategoria); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`UPDATE finanzas SET id_usuario = $1, id_movimiento_dinero = $2, id_categoria = $3, monto_presupuesto = $4, gasto = $5, disponible = $6, fecha = $7, activo = $8 WHERE id_finanzas = $9 RETURNING id_finanzas, id_usuario, id_movimiento_dinero, id_categoria, monto_presupuesto, gasto, disponible, fecha, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdMovimientoDinero, item.IdCategoria, item.MontoPresupuesto, item.Gasto, item.Disponible, item.Fecha, item.Activo, id)
	item, err = scanFinanzas(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "finanzas no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar finanzas")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarFinanzas(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "finanzas", "id_finanzas", "finanzas")
}

func scanFinanzas(scan func(dest ...any) error) (models.Finanzas, error) {
	var item models.Finanzas
	var idMovimiento sql.NullInt32
	var idCategoria sql.NullInt32
	var montoPresupuesto sql.NullFloat64
	var gasto sql.NullFloat64
	var disponible sql.NullFloat64
	var fecha sql.NullTime
	err := scan(&item.IdFinanzas, &item.IdUsuario, &idMovimiento, &idCategoria, &montoPresupuesto, &gasto, &disponible, &fecha, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.Finanzas{}, err
	}
	item.IdMovimientoDinero = nullInt32ToPointer(idMovimiento)
	item.IdCategoria = nullInt32ToPointer(idCategoria)
	item.MontoPresupuesto = nullFloat64ToPointer(montoPresupuesto)
	item.Gasto = nullFloat64ToPointer(gasto)
	item.Disponible = nullFloat64ToPointer(disponible)
	item.Fecha = nullTimeToPointer(fecha)
	return item, nil
}

func validarRelacionesFinanzas(idUsuario int, idMovimiento *int, idCategoria *int) error {
	existeUsuario, err := existeRegistro("auth.usuario", "id_usuario", idUsuario)
	if err != nil {
		return errors.New("error al validar id_usuario")
	}
	if !existeUsuario {
		return errors.New("id_usuario no existe")
	}
	if idMovimiento != nil {
		existe, err := existeRegistro("movimiento_ingreso_egreso", "id_movimiento_dinero", *idMovimiento)
		if err != nil {
			return errors.New("error al validar id_movimiento_dinero")
		}
		if !existe {
			return errors.New("id_movimiento_dinero no existe")
		}
	}
	if idCategoria != nil {
		existe, err := existeRegistro("categoria", "id_categoria", *idCategoria)
		if err != nil {
			return errors.New("error al validar id_categoria")
		}
		if !existe {
			return errors.New("id_categoria no existe")
		}
	}
	return nil
}
