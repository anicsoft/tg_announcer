package service

/*type serv struct {
	imageRepo repository.ImageRepository
	txManager db.TxManager
}

func New(
	imageRepo repository.ImageRepository,
	txManager db.TxManager,
) ImageService {
	return &serv{
		imageRepo: imageRepo,
		txManager: txManager,
	}
}

func (s serv) Get(ctx context.Context, parentId int) ([]string, error) {
	paths, err := s.imageRepo.Get(ctx, parentId)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (s serv) Upload(ctx context.Context, parentId int, paths []string) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		for _, path := range paths {
			_, err := s.imageRepo.Add(ctx, parentId, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}*/
