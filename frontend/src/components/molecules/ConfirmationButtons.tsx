'use client';

import Button from '../atoms/Button';

interface ConfirmationButtonsProps {
  onCancel: () => void;
  onConfirm: () => void;
  cancelText?: string;
  confirmText?: string;
  confirmVariant?: 'primary' | 'danger';
  className?: string;
}

export default function ConfirmationButtons({
  onCancel,
  onConfirm,
  cancelText = 'キャンセル',
  confirmText = '確認',
  confirmVariant = 'primary',
  className = '',
}: ConfirmationButtonsProps) {
  return (
    <div className={`flex justify-end space-x-3 ${className}`}>
      <Button
        variant="secondary"
        onClick={onCancel}
      >
        {cancelText}
      </Button>
      <Button
        variant={confirmVariant}
        onClick={onConfirm}
      >
        {confirmText}
      </Button>
    </div>
  );
}