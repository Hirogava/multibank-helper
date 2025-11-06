import pandas as pd
import numpy as np
import joblib
import os
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_absolute_error

# -----------------------------
# 1. Генерация синтетических транзакций
# -----------------------------

def generate_synthetic_transactions(n_users=50, months=12, seed=42):
    np.random.seed(seed)
    transactions = []

    categories = [
        "💼 Зарплата",
        "💰 Подработка/Бонус",
        "🏪 Продукты",
        "🎬 Развлечения/Покупки",
        "🚌 Транспорт",
        "🏠 ЖКХ/Аренда",
    ]

    for user_id in range(n_users):
        base_salary = np.random.randint(80000, 150000)
        for month in range(months):
            date = pd.Timestamp("2024-01-01") + pd.DateOffset(months=month)

            # Доходы
            income_salary = base_salary + np.random.randint(-10000, 10000)
            income_bonus = np.random.randint(0, 20000) if np.random.rand() < 0.3 else 0

            # Расходы
            rent = np.random.randint(20000, 30000)
            food = np.random.randint(8000, 15000)
            transport = np.random.randint(1500, 4000)
            entertainment = np.random.randint(1000, 6000)

            income = income_salary + income_bonus
            expenses = rent + food + transport + entertainment
            balance_change = income - expenses

            transactions.append({
                "user_id": user_id,
                "month": date.strftime("%Y-%m"),
                "income": income,
                "expenses": expenses,
                "transactions": np.random.randint(20, 80),
                "balance_change": balance_change,
            })

    df = pd.DataFrame(transactions)
    df["savings_rate"] = df["balance_change"] / (df["income"] + 1e-6)
    df["target_savings"] = df["balance_change"].shift(-1)  # прогноз на след. месяц
    df = df.dropna()

    return df


# -----------------------------
# 2. Обучение модели
# -----------------------------

def train_model(df, model_path="model_savings.pkl"):
    X = df[["income", "expenses", "transactions", "savings_rate"]]
    y = df["target_savings"]

    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.2, random_state=42
    )

    model = RandomForestRegressor(n_estimators=300, random_state=42)
    model.fit(X_train, y_train)

    preds = model.predict(X_test)
    mae = mean_absolute_error(y_test, preds)

    print(f"✅ Модель обучена. MAE = {mae:.2f}")

    # Убедимся, что папка для сохранения модели существует и составим путь к файлу
    output_dir = os.path.join(os.path.dirname(__file__), 'Ready_models')
    os.makedirs(output_dir, exist_ok=True)

    # если model_path передан как относительный файл, сохраняем его внутрь Ready_models
    if os.path.isabs(model_path):
        full_path = model_path
    else:
        full_path = os.path.join(output_dir, model_path)

    # Сохраняем модель в указанный файл
    joblib.dump(model, full_path)
    print(f"📦 Модель сохранена в {full_path}")

    return model


# -----------------------------
# 3. Запуск
# -----------------------------

if __name__ == "__main__":
    df = generate_synthetic_transactions(n_users=100, months=12)
    print("📊 Пример данных:")
    print(df.head())

    model = train_model(df)
