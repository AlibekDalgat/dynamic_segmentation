package handler

import "time"

func (h *Handler) DeleteExpirated() {
	go func() {
		for {
			h.services.DeleteExpirated()
			time.Sleep(5 * time.Second)
		}
	}()
}
